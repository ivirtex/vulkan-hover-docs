# VkCommandBufferInheritanceInfo(3) Manual Page

## Name

VkCommandBufferInheritanceInfo - Structure specifying command buffer inheritance information



## [](#_c_specification)C Specification

If the command buffer is a secondary command buffer, then the `VkCommandBufferInheritanceInfo` structure defines any state that will be inherited from the primary command buffer:

```c++
// Provided by VK_VERSION_1_0
typedef struct VkCommandBufferInheritanceInfo {
    VkStructureType                  sType;
    const void*                      pNext;
    VkRenderPass                     renderPass;
    uint32_t                         subpass;
    VkFramebuffer                    framebuffer;
    VkBool32                         occlusionQueryEnable;
    VkQueryControlFlags              queryFlags;
    VkQueryPipelineStatisticFlags    pipelineStatistics;
} VkCommandBufferInheritanceInfo;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `renderPass` is a [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html) object defining which render passes the `VkCommandBuffer` will be [compatible](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#renderpass-compatibility) with and **can** be executed within.
- `subpass` is the index of the subpass within the render pass instance that the `VkCommandBuffer` will be executed within.
- `framebuffer` **can** refer to the [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html) object that the `VkCommandBuffer` will be rendering to if it is executed within a render pass instance. It **can** be [VK\_NULL\_HANDLE](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NULL_HANDLE.html) if the framebuffer is not known.
  
  Note
  
  Specifying the exact framebuffer that the secondary command buffer will be executed with **may** result in better performance at command buffer execution time.
- `occlusionQueryEnable` specifies whether the command buffer **can** be executed while an occlusion query is active in the primary command buffer. If this is `VK_TRUE`, then this command buffer **can** be executed whether the primary command buffer has an occlusion query active or not. If this is `VK_FALSE`, then the primary command buffer **must** not have an occlusion query active.
- `queryFlags` specifies the query flags that **can** be used by an active occlusion query in the primary command buffer when this secondary command buffer is executed. If this value includes the `VK_QUERY_CONTROL_PRECISE_BIT` bit, then the active query **can** return boolean results or actual sample counts. If this bit is not set, then the active query **must** not use the `VK_QUERY_CONTROL_PRECISE_BIT` bit.
- `pipelineStatistics` is a bitmask of [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPipelineStatisticFlagBits.html) specifying the set of pipeline statistics that **can** be counted by an active query in the primary command buffer when this secondary command buffer is executed. If this value includes a given bit, then this command buffer **can** be executed whether the primary command buffer has a pipeline statistics query active that includes this bit or not. If this value excludes a given bit, then the active pipeline statistics query **must** not be from a query pool that counts that statistic.

## [](#_description)Description

If the [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBuffer.html) will not be executed within a render pass instance, or if the render pass instance was begun with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdBeginRendering.html), `renderPass`, `subpass`, and `framebuffer` are ignored.

Valid Usage

- [](#VUID-VkCommandBufferInheritanceInfo-occlusionQueryEnable-00056)VUID-VkCommandBufferInheritanceInfo-occlusionQueryEnable-00056  
  If the [`inheritedQueries`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-inheritedQueries) feature is not enabled, `occlusionQueryEnable` **must** be `VK_FALSE`
- [](#VUID-VkCommandBufferInheritanceInfo-queryFlags-00057)VUID-VkCommandBufferInheritanceInfo-queryFlags-00057  
  If the [`inheritedQueries`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-inheritedQueries) feature is enabled, `queryFlags` **must** be a valid combination of [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlagBits.html) values
- [](#VUID-VkCommandBufferInheritanceInfo-queryFlags-02788)VUID-VkCommandBufferInheritanceInfo-queryFlags-02788  
  If the [`inheritedQueries`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-inheritedQueries) feature is not enabled, `queryFlags` **must** be `0`
- [](#VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-02789)VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-02789  
  If the [`pipelineStatisticsQuery`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineStatisticsQuery) feature is enabled, `pipelineStatistics` **must** be a valid combination of [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPipelineStatisticFlagBits.html) values
- [](#VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-00058)VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-00058  
  If the [`pipelineStatisticsQuery`](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#features-pipelineStatisticsQuery) feature is not enabled, `pipelineStatistics` **must** be `0`

Valid Usage (Implicit)

- [](#VUID-VkCommandBufferInheritanceInfo-sType-sType)VUID-VkCommandBufferInheritanceInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO`
- [](#VUID-VkCommandBufferInheritanceInfo-pNext-pNext)VUID-VkCommandBufferInheritanceInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the `pNext` chain **must** be either `NULL` or a pointer to a valid instance of [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAttachmentSampleCountInfoAMD.html), [VkCommandBufferInheritanceConditionalRenderingInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceConditionalRenderingInfoEXT.html), [VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html), [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceRenderingInfo.html), [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html), [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExternalFormatANDROID.html), [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMultiviewPerViewAttributesInfoNVX.html), [VkRenderPassTileShadingCreateInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPassTileShadingCreateInfoQCOM.html), [VkRenderingAttachmentLocationInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingAttachmentLocationInfo.html), [VkRenderingInputAttachmentIndexInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderingInputAttachmentIndexInfo.html), or [VkTileMemoryBindInfoQCOM](https://registry.khronos.org/vulkan/specs/latest/man/html/VkTileMemoryBindInfoQCOM.html)
- [](#VUID-VkCommandBufferInheritanceInfo-sType-unique)VUID-VkCommandBufferInheritanceInfo-sType-unique  
  The `sType` value of each structure in the `pNext` chain **must** be unique
- [](#VUID-VkCommandBufferInheritanceInfo-commonparent)VUID-VkCommandBufferInheritanceInfo-commonparent  
  Both of `framebuffer`, and `renderPass` that are valid handles of non-ignored parameters **must** have been created, allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html)

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCommandBufferBeginInfo.html), [VkFramebuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkFramebuffer.html), [VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryControlFlags.html), [VkQueryPipelineStatisticFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkQueryPipelineStatisticFlags.html), [VkRenderPass](https://registry.khronos.org/vulkan/specs/latest/man/html/VkRenderPass.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCommandBufferInheritanceInfo).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0