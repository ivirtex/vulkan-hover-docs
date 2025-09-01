# vkCmdBuildMicromapsEXT(3) Manual Page

## Name

vkCmdBuildMicromapsEXT - Build a micromap



## [](#_c_specification)C Specification

To build micromaps call:

```c++
// Provided by VK_EXT_opacity_micromap
void vkCmdBuildMicromapsEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    infoCount,
    const VkMicromapBuildInfoEXT*               pInfos);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `infoCount` is the number of micromaps to build. It specifies the number of the `pInfos` structures that **must** be provided.
- `pInfos` is a pointer to an array of `infoCount` [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structures defining the data used to build each micromap.

## [](#_description)Description

The `vkCmdBuildMicromapsEXT` command provides the ability to initiate multiple micromaps builds, however there is no ordering or synchronization implied between any of the individual micromap builds.

Note

This means that there **cannot** be any memory aliasing between any micromap memories or scratch memories being used by any of the builds.

Accesses to the micromap scratch buffers as identified by the [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)::`scratchData` buffer device addresses **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of (`VK_ACCESS_2_MICROMAP_READ_BIT_EXT` | `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`). Accesses to [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)::`dstMicromap` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_2_MICROMAP_WRITE_BIT_EXT`.

Accesses to other input buffers as identified by any used values of [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)::`data` or [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)::`triangleArray` **must** be [synchronized](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-dependencies) with the `VK_PIPELINE_STAGE_2_MICROMAP_BUILD_BIT_EXT` [pipeline stage](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-pipeline-stages) and an [access type](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#synchronization-access-types) of `VK_ACCESS_SHADER_READ_BIT`.

Valid Usage

- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07461)VUID-vkCmdBuildMicromapsEXT-pInfos-07461  
  For each `pInfos`\[i], `dstMicromap` **must** have been created with a value of [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)::`size` greater than or equal to the memory size required by the build operation, as returned by [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMicromapBuildSizesEXT.html) with `pBuildInfo` = `pInfos`\[i]
- [](#VUID-vkCmdBuildMicromapsEXT-mode-07462)VUID-vkCmdBuildMicromapsEXT-mode-07462  
  The `mode` member of each element of `pInfos` **must** be a valid [VkBuildMicromapModeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuildMicromapModeEXT.html) value
- [](#VUID-vkCmdBuildMicromapsEXT-dstMicromap-07463)VUID-vkCmdBuildMicromapsEXT-dstMicromap-07463  
  The `dstMicromap` member of any element of `pInfos` **must** be a valid [VkMicromapEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapEXT.html) handle
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07464)VUID-vkCmdBuildMicromapsEXT-pInfos-07464  
  For each element of `pInfos` its `type` member **must** match the value of [VkMicromapCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapCreateInfoEXT.html)::`type` when its `dstMicromap` was created
- [](#VUID-vkCmdBuildMicromapsEXT-dstMicromap-07465)VUID-vkCmdBuildMicromapsEXT-dstMicromap-07465  
  The range of memory backing the `dstMicromap` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `dstMicromap` member of any other element of `pInfos`, which is accessed by this command
- [](#VUID-vkCmdBuildMicromapsEXT-dstMicromap-07466)VUID-vkCmdBuildMicromapsEXT-dstMicromap-07466  
  The range of memory backing the `dstMicromap` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any element of `pInfos` (including the same element), which is accessed by this command
- [](#VUID-vkCmdBuildMicromapsEXT-scratchData-07467)VUID-vkCmdBuildMicromapsEXT-scratchData-07467  
  The range of memory backing the `scratchData` member of any element of `pInfos` that is accessed by this command **must** not overlap the memory backing the `scratchData` member of any other element of `pInfos`, which is accessed by this command

<!--THE END-->

- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07508)VUID-vkCmdBuildMicromapsEXT-pInfos-07508  
  For each element of `pInfos`, the `buffer` used to create its `dstMicromap` member **must** be bound to device memory
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07509)VUID-vkCmdBuildMicromapsEXT-pInfos-07509  
  If `pInfos`\[i].`mode` is `VK_BUILD_MICROMAP_MODE_BUILD_EXT`, all addresses between `pInfos`\[i].`scratchData.deviceAddress` and `pInfos`\[i].`scratchData.deviceAddress` + N - 1 **must** be in the buffer device address range of the same buffer, where N is given by the `buildScratchSize` member of the [VkMicromapBuildSizesInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildSizesInfoEXT.html) structure returned from a call to [vkGetMicromapBuildSizesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetMicromapBuildSizesEXT.html) with an identical [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structure and primitive count
- [](#VUID-vkCmdBuildMicromapsEXT-data-07510)VUID-vkCmdBuildMicromapsEXT-data-07510  
  The buffers from which the buffer device addresses for all of the `data` and `triangleArray` members of all `pInfos`\[i] are queried **must** have been created with the `VK_BUFFER_USAGE_MICROMAP_BUILD_INPUT_READ_ONLY_BIT_EXT` usage flag
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07511)VUID-vkCmdBuildMicromapsEXT-pInfos-07511  
  For each element of `pInfos`\[i] the buffer from which the buffer device address `pInfos`\[i].`scratchData.deviceAddress` is queried **must** have been created with `VK_BUFFER_USAGE_STORAGE_BUFFER_BIT` usage flag
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07512)VUID-vkCmdBuildMicromapsEXT-pInfos-07512  
  For each element of `pInfos`, `scratchData.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-10896)VUID-vkCmdBuildMicromapsEXT-pInfos-10896  
  For each element of `pInfos`, `data.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-10897)VUID-vkCmdBuildMicromapsEXT-pInfos-10897  
  For each element of `pInfos`, `triangleArray.deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07514)VUID-vkCmdBuildMicromapsEXT-pInfos-07514  
  For each element of `pInfos`, its `scratchData.deviceAddress` member **must** be a multiple of [VkPhysicalDeviceAccelerationStructurePropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceAccelerationStructurePropertiesKHR.html)::`minAccelerationStructureScratchOffsetAlignment`
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-07515)VUID-vkCmdBuildMicromapsEXT-pInfos-07515  
  For each element of `pInfos`, its `triangleArray.deviceAddress` and `data.deviceAddress` members **must** be a multiple of `256`

Valid Usage (Implicit)

- [](#VUID-vkCmdBuildMicromapsEXT-commandBuffer-parameter)VUID-vkCmdBuildMicromapsEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdBuildMicromapsEXT-pInfos-parameter)VUID-vkCmdBuildMicromapsEXT-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html) structures
- [](#VUID-vkCmdBuildMicromapsEXT-commandBuffer-recording)VUID-vkCmdBuildMicromapsEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdBuildMicromapsEXT-commandBuffer-cmdpool)VUID-vkCmdBuildMicromapsEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support compute operations
- [](#VUID-vkCmdBuildMicromapsEXT-renderpass)VUID-vkCmdBuildMicromapsEXT-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdBuildMicromapsEXT-videocoding)VUID-vkCmdBuildMicromapsEXT-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdBuildMicromapsEXT-infoCount-arraylength)VUID-vkCmdBuildMicromapsEXT-infoCount-arraylength  
  `infoCount` **must** be greater than `0`

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

vkCmdBuildMicromapsEXT is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkMicromapBuildInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapBuildInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdBuildMicromapsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0