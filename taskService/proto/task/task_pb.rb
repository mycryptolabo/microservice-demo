# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: task/task.proto

require 'project/project_pb'
require 'user/user_pb'
require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("task/task.proto", :syntax => :proto3) do
    add_message "demo_task.ListTasksRequest" do
      optional :project_id, :string, 1
      optional :tag_id, :string, 2
      optional :assigned_user_id, :string, 3
    end
    add_message "demo_task.CreateTaskRequest" do
      optional :user_id, :string, 1
      optional :name, :string, 2
      optional :project_id, :string, 3
      optional :tag_id, :string, 4
      optional :assigned_user_id, :string, 5
    end
    add_message "demo_task.UpdateTaskRequest" do
      optional :user_id, :string, 1
      optional :task_id, :string, 2
      optional :tag_id, :string, 3
      optional :assigned_user_id, :string, 4
    end
    add_message "demo_task.ListTasksResponse" do
      repeated :tasks, :message, 1, "demo_task.TaskResponse"
    end
    add_message "demo_task.TaskResponse" do
      optional :id, :string, 1
      optional :user_id, :string, 2
      optional :name, :string, 3
      optional :project_id, :string, 4
      optional :tag_id, :string, 5
      optional :assigned_user_id, :string, 6
      optional :project, :message, 7, "demo_project.ProjectResponse"
      optional :tag, :message, 8, "demo_project.TagResponse"
      optional :assigned_user, :message, 9, "demo_user.VerifyResponse"
    end
  end
end

module DemoTask
  ListTasksRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("demo_task.ListTasksRequest").msgclass
  CreateTaskRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("demo_task.CreateTaskRequest").msgclass
  UpdateTaskRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("demo_task.UpdateTaskRequest").msgclass
  ListTasksResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("demo_task.ListTasksResponse").msgclass
  TaskResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("demo_task.TaskResponse").msgclass
end
