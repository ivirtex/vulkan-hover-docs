# vkCmdSetCoarseSampleOrderNV(3) Manual Page

## Name

vkCmdSetCoarseSampleOrderNV - Set order of coverage samples for coarse
fragments dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the order of coverage
samples in fragments larger than one pixel, call:

``` c
// Provided by VK_NV_shading_rate_image
void vkCmdSetCoarseSampleOrderNV(
    VkCommandBuffer                             commandBuffer,
    VkCoarseSampleOrderTypeNV                   sampleOrderType,
    uint32_t                                    customSampleOrderCount,
    const VkCoarseSampleOrderCustomNV*          pCustomSampleOrders);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `sampleOrderType` specifies the mechanism used to order coverage
  samples in fragments larger than one pixel.

- `customSampleOrderCount` specifies the number of custom sample
  orderings to use when ordering coverage samples.

- `pCustomSampleOrders` is a pointer to an array of
  [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html)
  structures, each structure specifying the coverage sample order for a
  single combination of fragment area and coverage sample count.

## <a href="#_description" class="anchor"></a>Description

If `sampleOrderType` is `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`, the
coverage sample order used for any combination of fragment area and
coverage sample count not enumerated in `pCustomSampleOrders` will be
identical to that used for `VK_COARSE_SAMPLE_ORDER_TYPE_DEFAULT_NV`.

This command sets the order of coverage samples for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with
`VK_DYNAMIC_STATE_VIEWPORT_COARSE_SAMPLE_ORDER_NV` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineViewportCoarseSampleOrderStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportCoarseSampleOrderStateCreateInfoNV.html)
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-02081"
  id="VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-02081"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-02081  
  If `sampleOrderType` is not `VK_COARSE_SAMPLE_ORDER_TYPE_CUSTOM_NV`,
  `customSamplerOrderCount` **must** be `0`

- <a href="#VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-02235"
  id="VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-02235"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-02235  
  The array `pCustomSampleOrders` **must** not contain two structures
  with matching values for both the `shadingRate` and `sampleCount`
  members

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-parameter"
  id="VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-parameter"
  id="VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-parameter"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-sampleOrderType-parameter  
  `sampleOrderType` **must** be a valid
  [VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderTypeNV.html) value

- <a
  href="#VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-parameter"
  id="VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-parameter"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-pCustomSampleOrders-parameter  
  If `customSampleOrderCount` is not `0`, `pCustomSampleOrders` **must**
  be a valid pointer to an array of `customSampleOrderCount` valid
  [VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html)
  structures

- <a href="#VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-recording"
  id="VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-recording"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetCoarseSampleOrderNV-videocoding"
  id="VUID-vkCmdSetCoarseSampleOrderNV-videocoding"></a>
  VUID-vkCmdSetCoarseSampleOrderNV-videocoding  
  This command **must** only be called outside of a video coding scope

Host Synchronization

- Host access to `commandBuffer` **must** be externally synchronized

- Host access to the `VkCommandPool` that `commandBuffer` was allocated
  from **must** be externally synchronized

Command Properties

<table class="tableblock frame-all grid-all stretch">
<colgroup>
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
<col style="width: 20%" />
</colgroup>
<thead>
<tr class="header">
<th class="tableblock halign-left valign-top"><a
href="#VkCommandBufferLevel">Command Buffer Levels</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginRenderPass">Render Pass Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#vkCmdBeginVideoCodingKHR">Video Coding Scope</a></th>
<th class="tableblock halign-left valign-top"><a
href="#VkQueueFlagBits">Supported Queue Types</a></th>
<th class="tableblock halign-left valign-top"><a
href="#fundamentals-queueoperation-command-types">Command Type</a></th>
</tr>
</thead>
<tbody>
<tr class="odd">
<td class="tableblock halign-left valign-top"><p>Primary<br />
Secondary</p></td>
<td class="tableblock halign-left valign-top"><p>Both</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_NV_shading_rate_image](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NV_shading_rate_image.html),
[VkCoarseSampleOrderCustomNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderCustomNV.html),
[VkCoarseSampleOrderTypeNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCoarseSampleOrderTypeNV.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetCoarseSampleOrderNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
