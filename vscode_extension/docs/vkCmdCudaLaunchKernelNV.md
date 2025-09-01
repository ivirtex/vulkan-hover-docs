# vkCmdCudaLaunchKernelNV(3) Manual Page

## Name

vkCmdCudaLaunchKernelNV - Dispatch compute work items



## [](#_c_specification)C Specification

To record a CUDA kernel launch, call:

```c++
// Provided by VK_NV_cuda_kernel_launch
void vkCmdCudaLaunchKernelNV(
    VkCommandBuffer                             commandBuffer,
    const VkCudaLaunchInfoNV*                   pLaunchInfo);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pLaunchInfo` is a pointer to a [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaLaunchInfoNV.html) structure in which the grid (similar to workgroup) dimension, function handle and related arguments are defined.

## [](#_description)Description

When the command is executed, a global workgroup consisting of `gridDimX` × `gridDimY` × `gridDimZ` local workgroups is assembled.

Valid Usage (Implicit)

- [](#VUID-vkCmdCudaLaunchKernelNV-commandBuffer-parameter)VUID-vkCmdCudaLaunchKernelNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCudaLaunchKernelNV-pLaunchInfo-parameter)VUID-vkCmdCudaLaunchKernelNV-pLaunchInfo-parameter  
  `pLaunchInfo` **must** be a valid pointer to a valid [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaLaunchInfoNV.html) structure
- [](#VUID-vkCmdCudaLaunchKernelNV-commandBuffer-recording)VUID-vkCmdCudaLaunchKernelNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCudaLaunchKernelNV-commandBuffer-cmdpool)VUID-vkCmdCudaLaunchKernelNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdCudaLaunchKernelNV-videocoding)VUID-vkCmdCudaLaunchKernelNV-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Both

Outside

Graphics  
Compute

Action

Conditional Rendering

vkCmdCudaLaunchKernelNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_cuda\_kernel\_launch](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cuda_kernel_launch.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCudaLaunchInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCudaLaunchInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCudaLaunchKernelNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0