# vkCmdConvertCooperativeVectorMatrixNV(3) Manual Page

## Name

vkCmdConvertCooperativeVectorMatrixNV - Convert a cooperative vector matrix from one layout and type to another



## [](#_c_specification)C Specification

To convert a matrix to another layout and type, call:

```c++
// Provided by VK_NV_cooperative_vector
void vkCmdConvertCooperativeVectorMatrixNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    infoCount,
    const VkConvertCooperativeVectorMatrixInfoNV* pInfos);
```

## [](#_parameters)Parameters

- `commandBuffer` is the command buffer into which the command will be recorded.
- `infoCount` is the number of layout conversions to perform.
- `pInfos` is a pointer to an array of [VkConvertCooperativeVectorMatrixInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConvertCooperativeVectorMatrixInfoNV.html) structures containing information about the layout conversion.

## [](#_description)Description

This command does the same conversions as [vkConvertCooperativeVectorMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkConvertCooperativeVectorMatrixNV.html), but executes on the device. One conversion is performed for each of the `infoCount` elements of `pInfos`.

This command’s execution is synchronized using `VK_PIPELINE_STAGE_2_CONVERT_COOPERATIVE_VECTOR_MATRIX_BIT_NV`.

Valid Usage

- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10083)VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10083  
  For each element of `pInfo`, `srcData`::`deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10895)VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10895  
  For each element of `pInfo`, `dstData`::`deviceAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html)
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10084)VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10084  
  For each element of `pInfo`, `srcData`::`deviceAddress` **must** be 64 byte aligned
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10085)VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10085  
  For each element of `pInfo`, `dstData`::`deviceAddress` **must** be 64 byte aligned
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10086)VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10086  
  For each element of `pInfo`, `srcSize` **must** be large enough to contain the source matrix, based either on the standard matrix layout or based on the size filled out by [vkConvertCooperativeVectorMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkConvertCooperativeVectorMatrixNV.html)
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10087)VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfo-10087  
  For each element of `pInfo`, the value pointed to by `pDstSize` **must** be large enough to contain the destination matrix, based either on the standard matrix layout or based on the size filled out by [vkConvertCooperativeVectorMatrixNV](https://registry.khronos.org/vulkan/specs/latest/man/html/vkConvertCooperativeVectorMatrixNV.html)
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-None-10088)VUID-vkCmdConvertCooperativeVectorMatrixNV-None-10088  
  Memory accessed by the sources and destinations of all of the conversions **must** not overlap

Valid Usage (Implicit)

- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-commandBuffer-parameter)VUID-vkCmdConvertCooperativeVectorMatrixNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) handle
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfos-parameter)VUID-vkCmdConvertCooperativeVectorMatrixNV-pInfos-parameter  
  `pInfos` **must** be a valid pointer to an array of `infoCount` valid [VkConvertCooperativeVectorMatrixInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConvertCooperativeVectorMatrixInfoNV.html) structures
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-commandBuffer-recording)VUID-vkCmdConvertCooperativeVectorMatrixNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording state](#commandbuffers-lifecycle)
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-commandBuffer-cmdpool)VUID-vkCmdConvertCooperativeVectorMatrixNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must** support graphics, or compute operations
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-renderpass)VUID-vkCmdConvertCooperativeVectorMatrixNV-renderpass  
  This command **must** only be called outside of a render pass instance
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-videocoding)VUID-vkCmdConvertCooperativeVectorMatrixNV-videocoding  
  This command **must** only be called outside of a video coding scope
- [](#VUID-vkCmdConvertCooperativeVectorMatrixNV-infoCount-arraylength)VUID-vkCmdConvertCooperativeVectorMatrixNV-infoCount-arraylength  
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

Graphics  
Compute

Action

Conditional Rendering

vkCmdConvertCooperativeVectorMatrixNV is not affected by [conditional rendering](#drawing-conditional-rendering)

## [](#_see_also)See Also

[VK\_NV\_cooperative\_vector](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_cooperative_vector.html), [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html), [VkConvertCooperativeVectorMatrixInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkConvertCooperativeVectorMatrixInfoNV.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkCmdConvertCooperativeVectorMatrixNV)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0