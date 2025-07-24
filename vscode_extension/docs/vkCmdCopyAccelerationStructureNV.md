# vkCmdCopyAccelerationStructureNV(3) Manual Page

## Name

vkCmdCopyAccelerationStructureNV - Copy an acceleration structure



## [](#_c_specification)C Specification

To copy an acceleration structure call:

```c++
// Provided by VK_NV_ray_tracing
void vkCmdCopyAccelerationStructureNV(
    VkCommandBuffer                             commandBuffer,
    VkAccelerationStructureNV                   dst,
    VkAccelerationStructureNV                   src,
    VkCopyAccelerationStructureModeKHR          mode);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `dst` is the target acceleration structure for the copy.
- `src` is the source acceleration structure for the copy.
- `mode` is a [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value specifying additional operations to perform during the copy.

## [](#_description)Description

Accesses to `src` and `dst` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_ACCELERATION_STRUCTURE_COPY_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) or the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages), and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR` as appropriate.

Valid Usage

- [](#VUID-vkCmdCopyAccelerationStructureNV-mode-03410)VUID-vkCmdCopyAccelerationStructureNV-mode-03410  
  `mode` **must** be `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR` or `VK_COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR`
- [](#VUID-vkCmdCopyAccelerationStructureNV-src-04963)VUID-vkCmdCopyAccelerationStructureNV-src-04963  
  The source acceleration structure `src` **must** have been constructed prior to the execution of this command
- [](#VUID-vkCmdCopyAccelerationStructureNV-src-03411)VUID-vkCmdCopyAccelerationStructureNV-src-03411  
  If `mode` is `VK_COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR`, `src` **must** have been constructed with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_COMPACTION_BIT_KHR` in the build
- [](#VUID-vkCmdCopyAccelerationStructureNV-buffer-03718)VUID-vkCmdCopyAccelerationStructureNV-buffer-03718  
  The `buffer` used to create `src` **must** be bound to device memory
- [](#VUID-vkCmdCopyAccelerationStructureNV-buffer-03719)VUID-vkCmdCopyAccelerationStructureNV-buffer-03719  
  The `buffer` used to create `dst` **must** be bound to device memory
- [](#VUID-vkCmdCopyAccelerationStructureNV-dst-07791)VUID-vkCmdCopyAccelerationStructureNV-dst-07791  
  The range of memory backing `dst` that is accessed by this command **must** not overlap the memory backing `src` that is accessed by this command
- [](#VUID-vkCmdCopyAccelerationStructureNV-dst-07792)VUID-vkCmdCopyAccelerationStructureNV-dst-07792  
  `dst` **must** be bound completely and contiguously to a single `VkDeviceMemory` object via [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindAccelerationStructureMemoryNV.html)

Valid Usage (Implicit)

- [](#VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-parameter)VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdCopyAccelerationStructureNV-dst-parameter)VUID-vkCmdCopyAccelerationStructureNV-dst-parameter  
  `dst` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-vkCmdCopyAccelerationStructureNV-src-parameter)VUID-vkCmdCopyAccelerationStructureNV-src-parameter  
  `src` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-vkCmdCopyAccelerationStructureNV-mode-parameter)VUID-vkCmdCopyAccelerationStructureNV-mode-parameter  
  `mode` **must** be a valid [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html) value
- [](#VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-recording)VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-cmdpool)VUID-vkCmdCopyAccelerationStructureNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdCopyAccelerationStructureNV-renderpass)VUID-vkCmdCopyAccelerationStructureNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdCopyAccelerationStructureNV-videocoding)VUID-vkCmdCopyAccelerationStructureNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdCopyAccelerationStructureNV-commonparent)VUID-vkCmdCopyAccelerationStructureNV-commonparent  
  Each of `commandBuffer`, `dst`, and `src` **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized
- Host access to the `VkCommandPool` that `commandBuffer` was allocated from **must** be externally synchronized

Command Properties

     [Command Buffer Levels](#VkCommandBufferLevel) [Render Pass Scope](#vkCmdBeginRenderPass) [Video Coding Scope](#vkCmdBeginVideoCodingKHR) [Supported Queue Types](#VkQueueFlagBits) [Command Type](#fundamentals-queueoperation-command-types)

Primary  
Secondary

Outside

Outside

Compute

Action

Conditional Rendering

vkCmdCopyAccelerationStructureNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkCopyAccelerationStructureModeKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyAccelerationStructureModeKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdCopyAccelerationStructureNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0