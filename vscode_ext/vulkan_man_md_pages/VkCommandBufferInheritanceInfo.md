# VkCommandBufferInheritanceInfo(3) Manual Page

## Name

VkCommandBufferInheritanceInfo - Structure specifying command buffer
inheritance information



## <a href="#_c_specification" class="anchor"></a>C Specification

If the command buffer is a secondary command buffer, then the
`VkCommandBufferInheritanceInfo` structure defines any state that will
be inherited from the primary command buffer:

``` c
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

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `renderPass` is a [VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html) object defining
  which render passes the `VkCommandBuffer` will be <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#renderpass-compatibility"
  target="_blank" rel="noopener">compatible</a> with and **can** be
  executed within.

- `subpass` is the index of the subpass within the render pass instance
  that the `VkCommandBuffer` will be executed within.

- `framebuffer` **can** refer to the [VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html)
  object that the `VkCommandBuffer` will be rendering to if it is
  executed within a render pass instance. It **can** be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) if the framebuffer is not known.

  <table>
  <colgroup>
  <col style="width: 50%" />
  <col style="width: 50%" />
  </colgroup>
  <tbody>
  <tr>
  <td class="icon"><em></em></td>
  <td class="content">Note
  <p>Specifying the exact framebuffer that the secondary command buffer
  will be executed with <strong>may</strong> result in better performance
  at command buffer execution time.</p></td>
  </tr>
  </tbody>
  </table>

- `occlusionQueryEnable` specifies whether the command buffer **can** be
  executed while an occlusion query is active in the primary command
  buffer. If this is `VK_TRUE`, then this command buffer **can** be
  executed whether the primary command buffer has an occlusion query
  active or not. If this is `VK_FALSE`, then the primary command buffer
  **must** not have an occlusion query active.

- `queryFlags` specifies the query flags that **can** be used by an
  active occlusion query in the primary command buffer when this
  secondary command buffer is executed. If this value includes the
  `VK_QUERY_CONTROL_PRECISE_BIT` bit, then the active query **can**
  return boolean results or actual sample counts. If this bit is not
  set, then the active query **must** not use the
  `VK_QUERY_CONTROL_PRECISE_BIT` bit.

- `pipelineStatistics` is a bitmask of
  [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html)
  specifying the set of pipeline statistics that **can** be counted by
  an active query in the primary command buffer when this secondary
  command buffer is executed. If this value includes a given bit, then
  this command buffer **can** be executed whether the primary command
  buffer has a pipeline statistics query active that includes this bit
  or not. If this value excludes a given bit, then the active pipeline
  statistics query **must** not be from a query pool that counts that
  statistic.

## <a href="#_description" class="anchor"></a>Description

If the [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) will not be executed
within a render pass instance, or if the render pass instance was begun
with [vkCmdBeginRendering](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBeginRendering.html), `renderPass`,
`subpass`, and `framebuffer` are ignored.

Valid Usage

- <a
  href="#VUID-VkCommandBufferInheritanceInfo-occlusionQueryEnable-00056"
  id="VUID-VkCommandBufferInheritanceInfo-occlusionQueryEnable-00056"></a>
  VUID-VkCommandBufferInheritanceInfo-occlusionQueryEnable-00056  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-inheritedQueries"
  target="_blank" rel="noopener"><code>inheritedQueries</code></a>
  feature is not enabled, `occlusionQueryEnable` **must** be `VK_FALSE`

- <a href="#VUID-VkCommandBufferInheritanceInfo-queryFlags-00057"
  id="VUID-VkCommandBufferInheritanceInfo-queryFlags-00057"></a>
  VUID-VkCommandBufferInheritanceInfo-queryFlags-00057  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-inheritedQueries"
  target="_blank" rel="noopener"><code>inheritedQueries</code></a>
  feature is enabled, `queryFlags` **must** be a valid combination of
  [VkQueryControlFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlagBits.html) values

- <a href="#VUID-VkCommandBufferInheritanceInfo-queryFlags-02788"
  id="VUID-VkCommandBufferInheritanceInfo-queryFlags-02788"></a>
  VUID-VkCommandBufferInheritanceInfo-queryFlags-02788  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-inheritedQueries"
  target="_blank" rel="noopener"><code>inheritedQueries</code></a>
  feature is not enabled, `queryFlags` **must** be `0`

- <a href="#VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-02789"
  id="VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-02789"></a>
  VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-02789  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineStatisticsQuery"
  target="_blank" rel="noopener"><code>pipelineStatisticsQuery</code></a>
  feature is enabled, `pipelineStatistics` **must** be a valid
  combination of
  [VkQueryPipelineStatisticFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlagBits.html)
  values

- <a href="#VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-00058"
  id="VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-00058"></a>
  VUID-VkCommandBufferInheritanceInfo-pipelineStatistics-00058  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-pipelineStatisticsQuery"
  target="_blank" rel="noopener"><code>pipelineStatisticsQuery</code></a>
  feature is not enabled, `pipelineStatistics` **must** be `0`

Valid Usage (Implicit)

- <a href="#VUID-VkCommandBufferInheritanceInfo-sType-sType"
  id="VUID-VkCommandBufferInheritanceInfo-sType-sType"></a>
  VUID-VkCommandBufferInheritanceInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_COMMAND_BUFFER_INHERITANCE_INFO`

- <a href="#VUID-VkCommandBufferInheritanceInfo-pNext-pNext"
  id="VUID-VkCommandBufferInheritanceInfo-pNext-pNext"></a>
  VUID-VkCommandBufferInheritanceInfo-pNext-pNext  
  Each `pNext` member of any structure (including this one) in the
  `pNext` chain **must** be either `NULL` or a pointer to a valid
  instance of
  [VkAttachmentSampleCountInfoAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAttachmentSampleCountInfoAMD.html),
  [VkCommandBufferInheritanceConditionalRenderingInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceConditionalRenderingInfoEXT.html),
  [VkCommandBufferInheritanceRenderPassTransformInfoQCOM](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderPassTransformInfoQCOM.html),
  [VkCommandBufferInheritanceRenderingInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceRenderingInfo.html),
  [VkCommandBufferInheritanceViewportScissorInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferInheritanceViewportScissorInfoNV.html),
  [VkExternalFormatANDROID](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExternalFormatANDROID.html),
  [VkMultiviewPerViewAttributesInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMultiviewPerViewAttributesInfoNVX.html),
  [VkRenderingAttachmentLocationInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingAttachmentLocationInfoKHR.html),
  or
  [VkRenderingInputAttachmentIndexInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderingInputAttachmentIndexInfoKHR.html)

- <a href="#VUID-VkCommandBufferInheritanceInfo-sType-unique"
  id="VUID-VkCommandBufferInheritanceInfo-sType-unique"></a>
  VUID-VkCommandBufferInheritanceInfo-sType-unique  
  The `sType` value of each struct in the `pNext` chain **must** be
  unique

- <a href="#VUID-VkCommandBufferInheritanceInfo-commonparent"
  id="VUID-VkCommandBufferInheritanceInfo-commonparent"></a>
  VUID-VkCommandBufferInheritanceInfo-commonparent  
  Both of `framebuffer`, and `renderPass` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html),
[VkCommandBufferBeginInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferBeginInfo.html),
[VkFramebuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkFramebuffer.html),
[VkQueryControlFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryControlFlags.html),
[VkQueryPipelineStatisticFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkQueryPipelineStatisticFlags.html),
[VkRenderPass](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPass.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCommandBufferInheritanceInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
