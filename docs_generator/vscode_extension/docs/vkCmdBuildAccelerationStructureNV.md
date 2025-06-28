# vkCmdBuildAccelerationStructureNV(3) Manual Page

## Name

vkCmdBuildAccelerationStructureNV - Build an acceleration structure



## [](#_c_specification)C Specification

To build an acceleration structure call:

```c++
// Provided by VK_NV_ray_tracing
void vkCmdBuildAccelerationStructureNV(
    VkCommandBuffer                             commandBuffer,
    const VkAccelerationStructureInfoNV*        pInfo,
    VkBuffer                                    instanceData,
    VkDeviceSize                                instanceOffset,
    VkBool32                                    update,
    VkAccelerationStructureNV                   dst,
    VkAccelerationStructureNV                   src,
    VkBuffer                                    scratch,
    VkDeviceSize                                scratchOffset);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `pInfo` contains the shared information for the acceleration structure’s structure.
- `instanceData` is the buffer containing an array of [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html) structures defining acceleration structures. This parameter **must** be `NULL` for bottom level acceleration structures.
- `instanceOffset` is the offset in bytes (relative to the start of `instanceData`) at which the instance data is located.
- `update` specifies whether to update the `dst` acceleration structure with the data in `src`.
- `dst` is a pointer to the target acceleration structure for the build.
- `src` is a pointer to an existing acceleration structure that is to be used to update the `dst` acceleration structure.
- `scratch` is the [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) that will be used as scratch memory for the build.
- `scratchOffset` is the offset in bytes relative to the start of `scratch` that will be used as a scratch memory.

## [](#_description)Description

Accesses to `dst`, `src`, and `scratch` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_ACCELERATION_STRUCTURE_BUILD_BIT_KHR` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_ACCELERATION_STRUCTURE_READ_BIT_KHR` or `VK_ACCESS_ACCELERATION_STRUCTURE_WRITE_BIT_KHR`.

Valid Usage

- [](#VUID-vkCmdBuildAccelerationStructureNV-geometryCount-02241)VUID-vkCmdBuildAccelerationStructureNV-geometryCount-02241  
  `geometryCount` **must** be less than or equal to [VkPhysicalDeviceRayTracingPropertiesNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceRayTracingPropertiesNV.html)::`maxGeometryCount`
- [](#VUID-vkCmdBuildAccelerationStructureNV-dst-02488)VUID-vkCmdBuildAccelerationStructureNV-dst-02488  
  `dst` **must** have been created with compatible [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html) where [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)::`type` and [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)::`flags` are identical, [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)::`instanceCount` and [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)::`geometryCount` for `dst` are greater than or equal to the build size and each geometry in [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)::`pGeometries` for `dst` has greater than or equal to the number of vertices, indices, and AABBs
- [](#VUID-vkCmdBuildAccelerationStructureNV-update-02489)VUID-vkCmdBuildAccelerationStructureNV-update-02489  
  If `update` is `VK_TRUE`, `src` **must** not be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html)
- [](#VUID-vkCmdBuildAccelerationStructureNV-update-02490)VUID-vkCmdBuildAccelerationStructureNV-update-02490  
  If `update` is `VK_TRUE`, `src` **must** have previously been constructed with `VK_BUILD_ACCELERATION_STRUCTURE_ALLOW_UPDATE_BIT_NV` set in [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html)::`flags` in the original build
- [](#VUID-vkCmdBuildAccelerationStructureNV-update-02491)VUID-vkCmdBuildAccelerationStructureNV-update-02491  
  If `update` is `VK_FALSE`, the `size` member of the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure returned from a call to [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html) with [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`accelerationStructure` set to `dst` and [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`type` set to `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_BUILD_SCRATCH_NV` **must** be less than or equal to the size of `scratch` minus `scratchOffset`
- [](#VUID-vkCmdBuildAccelerationStructureNV-update-02492)VUID-vkCmdBuildAccelerationStructureNV-update-02492  
  If `update` is `VK_TRUE`, the `size` member of the [VkMemoryRequirements](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMemoryRequirements.html) structure returned from a call to [vkGetAccelerationStructureMemoryRequirementsNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureMemoryRequirementsNV.html) with [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`accelerationStructure` set to `dst` and [VkAccelerationStructureMemoryRequirementsInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureMemoryRequirementsInfoNV.html)::`type` set to `VK_ACCELERATION_STRUCTURE_MEMORY_REQUIREMENTS_TYPE_UPDATE_SCRATCH_NV` **must** be less than or equal to the size of `scratch` minus `scratchOffset`
- [](#VUID-vkCmdBuildAccelerationStructureNV-scratch-03522)VUID-vkCmdBuildAccelerationStructureNV-scratch-03522  
  `scratch` **must** have been created with `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` usage flag
- [](#VUID-vkCmdBuildAccelerationStructureNV-instanceData-03523)VUID-vkCmdBuildAccelerationStructureNV-instanceData-03523  
  If `instanceData` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `instanceData` **must** have been created with `VK_BUFFER_USAGE_RAY_TRACING_BIT_NV` usage flag
- [](#VUID-vkCmdBuildAccelerationStructureNV-accelerationStructureReference-03786)VUID-vkCmdBuildAccelerationStructureNV-accelerationStructureReference-03786  
  Each [VkAccelerationStructureInstanceKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInstanceKHR.html)::`accelerationStructureReference` value in `instanceData` **must** be a valid device address containing a value obtained from [vkGetAccelerationStructureHandleNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetAccelerationStructureHandleNV.html)
- [](#VUID-vkCmdBuildAccelerationStructureNV-update-03524)VUID-vkCmdBuildAccelerationStructureNV-update-03524  
  If `update` is `VK_TRUE`, then objects that were previously active **must** not be made inactive as per [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-inactive-prims](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-inactive-prims)
- [](#VUID-vkCmdBuildAccelerationStructureNV-update-03525)VUID-vkCmdBuildAccelerationStructureNV-update-03525  
  If `update` is `VK_TRUE`, then objects that were previously inactive **must** not be made active as per [https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-inactive-prims](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#acceleration-structure-inactive-prims)
- [](#VUID-vkCmdBuildAccelerationStructureNV-update-03526)VUID-vkCmdBuildAccelerationStructureNV-update-03526  
  If `update` is `VK_TRUE`, the `src` and `dst` objects **must** either be the same object or not have any [memory aliasing](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#resources-memory-aliasing)
- [](#VUID-vkCmdBuildAccelerationStructureNV-dst-07787)VUID-vkCmdBuildAccelerationStructureNV-dst-07787  
  `dst` **must** be bound completely and contiguously to a single `VkDeviceMemory` object via [vkBindAccelerationStructureMemoryNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkBindAccelerationStructureMemoryNV.html)

Valid Usage (Implicit)

- [](#VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-parameter)VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBuildAccelerationStructureNV-pInfo-parameter)VUID-vkCmdBuildAccelerationStructureNV-pInfo-parameter  
  `pInfo` **must** be a valid pointer to a valid [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html) structure
- [](#VUID-vkCmdBuildAccelerationStructureNV-instanceData-parameter)VUID-vkCmdBuildAccelerationStructureNV-instanceData-parameter  
  If `instanceData` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `instanceData` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdBuildAccelerationStructureNV-dst-parameter)VUID-vkCmdBuildAccelerationStructureNV-dst-parameter  
  `dst` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-vkCmdBuildAccelerationStructureNV-src-parameter)VUID-vkCmdBuildAccelerationStructureNV-src-parameter  
  If `src` is not [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html), `src` **must** be a valid [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html) handle
- [](#VUID-vkCmdBuildAccelerationStructureNV-scratch-parameter)VUID-vkCmdBuildAccelerationStructureNV-scratch-parameter  
  `scratch` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle
- [](#VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-recording)VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-cmdpool)VUID-vkCmdBuildAccelerationStructureNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBuildAccelerationStructureNV-renderpass)VUID-vkCmdBuildAccelerationStructureNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBuildAccelerationStructureNV-videocoding)VUID-vkCmdBuildAccelerationStructureNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBuildAccelerationStructureNV-commonparent)VUID-vkCmdBuildAccelerationStructureNV-commonparent  
  Each of `commandBuffer`, `dst`, `instanceData`, `scratch`, and `src` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

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

## [](#_see_also)See Also

[VK\_NV\_ray\_tracing](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_ray_tracing.html), [VkAccelerationStructureInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureInfoNV.html), [VkAccelerationStructureNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAccelerationStructureNV.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBuildAccelerationStructureNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0