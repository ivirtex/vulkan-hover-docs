# vkCmdInitializeGraphScratchMemoryAMDX(3) Manual Page

## Name

vkCmdInitializeGraphScratchMemoryAMDX - Initialize scratch memory for an execution graph



## [](#_c_specification)C Specification

To initialize scratch memory for a particular execution graph, call:

```c++
// Provided by VK_AMDX_shader_enqueue
void vkCmdInitializeGraphScratchMemoryAMDX(
    VkCommandBuffer                             commandBuffer,
    VkPipeline                                  executionGraph,
    VkDeviceAddress                             scratch,
    VkDeviceSize                                scratchSize);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `executionGraph` is the execution graph pipeline to initialize the scratch memory for.
- `scratch` is the address of scratch memory to be initialized.
- `scratchSize` is a range in bytes of scratch memory to be initialized.

## [](#_description)Description

This command **must** be called before using `scratch` to dispatch the bound execution graph pipeline.

Execution of this command **may** modify any memory locations in the range [`scratch`,`scratch` + `scratchSize`). Accesses to this memory range are performed in the `VK_PIPELINE_STAGE_2_COMPUTE_SHADER_BIT` pipeline stage with the `VK_ACCESS_2_SHADER_STORAGE_READ_BIT` and `VK_ACCESS_2_SHADER_STORAGE_WRITE_BIT` access flags.

If any portion of `scratch` is modified by any command other than [vkCmdDispatchGraphAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphAMDX.html), [vkCmdDispatchGraphIndirectAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphIndirectAMDX.html), [vkCmdDispatchGraphIndirectCountAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdDispatchGraphIndirectCountAMDX.html), or [vkCmdInitializeGraphScratchMemoryAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdInitializeGraphScratchMemoryAMDX.html) with the same execution graph, it **must** be reinitialized for the execution graph again before dispatching against it.

Valid Usage

- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-10185)VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-10185  
  `scratch` **must** be the device address of an allocated memory range at least as large as `scratchSize`
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratchSize-10186)VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratchSize-10186  
  `scratchSize` **must** be greater than or equal to [VkExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExecutionGraphPipelineScratchSizeAMDX.html)::`minSize` returned by [vkGetExecutionGraphPipelineScratchSizeAMDX](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetExecutionGraphPipelineScratchSizeAMDX.html) for the bound execution graph pipeline
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09144)VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-09144  
  `scratch` **must** be a multiple of 64

Valid Usage (Implicit)

- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-parameter)VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-executionGraph-parameter)VUID-vkCmdInitializeGraphScratchMemoryAMDX-executionGraph-parameter  
  `executionGraph` **must** be a valid [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html) handle
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-parameter)VUID-vkCmdInitializeGraphScratchMemoryAMDX-scratch-parameter  
  `scratch` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-recording)VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-cmdpool)VUID-vkCmdInitializeGraphScratchMemoryAMDX-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support VK\_QUEUE\_COMPUTE\_BIT, or VK\_QUEUE\_GRAPHICS\_BIT operations
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-videocoding)VUID-vkCmdInitializeGraphScratchMemoryAMDX-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-bufferlevel)VUID-vkCmdInitializeGraphScratchMemoryAMDX-bufferlevel  
  `commandBuffer` **must** be a primary `VkCommandBuffer`
- [](#VUID-vkCmdInitializeGraphScratchMemoryAMDX-commonparent)VUID-vkCmdInitializeGraphScratchMemoryAMDX-commonparent  
  Both of `commandBuffer`, and `executionGraph` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary

Both

Outside

VK\_QUEUE\_COMPUTE\_BIT  
VK\_QUEUE\_GRAPHICS\_BIT

Action

Conditional Rendering

vkCmdInitializeGraphScratchMemoryAMDX is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_AMDX\_shader\_enqueue](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMDX_shader_enqueue.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html), [VkPipeline](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPipeline.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdInitializeGraphScratchMemoryAMDX).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0