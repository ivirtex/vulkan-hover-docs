# vkCmdSetSampleLocationsEXT(3) Manual Page

## Name

vkCmdSetSampleLocationsEXT - Set sample locations dynamically for a
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the sample locations
used for rasterization, call:

``` c
// Provided by VK_EXT_sample_locations
void vkCmdSetSampleLocationsEXT(
    VkCommandBuffer                             commandBuffer,
    const VkSampleLocationsInfoEXT*             pSampleLocationsInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pSampleLocationsInfo` is the sample locations state to set.

## <a href="#_description" class="anchor"></a>Description

This command sets the custom sample locations for subsequent drawing
commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with `VK_DYNAMIC_STATE_SAMPLE_LOCATIONS_EXT` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`,
and when the
[VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsEnable`
property of the bound graphics pipeline is `VK_TRUE`. Otherwise, this
state is specified by the
[VkPipelineSampleLocationsStateCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineSampleLocationsStateCreateInfoEXT.html)::`sampleLocationsInfo`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetSampleLocationsEXT-variableSampleLocations-01530"
  id="VUID-vkCmdSetSampleLocationsEXT-variableSampleLocations-01530"></a>
  VUID-vkCmdSetSampleLocationsEXT-variableSampleLocations-01530  
  If
  [VkPhysicalDeviceSampleLocationsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceSampleLocationsPropertiesEXT.html)::`variableSampleLocations`
  is `VK_FALSE` then the current render pass **must** have been begun by
  specifying a
  [VkRenderPassSampleLocationsBeginInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkRenderPassSampleLocationsBeginInfoEXT.html)
  structure whose `pPostSubpassSampleLocations` member contains an
  element with a `subpassIndex` matching the current subpass index and
  the `sampleLocationsInfo` member of that element **must** match the
  sample locations state pointed to by `pSampleLocationsInfo`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdSetSampleLocationsEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetSampleLocationsEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetSampleLocationsEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetSampleLocationsEXT-pSampleLocationsInfo-parameter"
  id="VUID-vkCmdSetSampleLocationsEXT-pSampleLocationsInfo-parameter"></a>
  VUID-vkCmdSetSampleLocationsEXT-pSampleLocationsInfo-parameter  
  `pSampleLocationsInfo` **must** be a valid pointer to a valid
  [VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html) structure

- <a href="#VUID-vkCmdSetSampleLocationsEXT-commandBuffer-recording"
  id="VUID-vkCmdSetSampleLocationsEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetSampleLocationsEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetSampleLocationsEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetSampleLocationsEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetSampleLocationsEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetSampleLocationsEXT-videocoding"
  id="VUID-vkCmdSetSampleLocationsEXT-videocoding"></a>
  VUID-vkCmdSetSampleLocationsEXT-videocoding  
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
<tr>
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
<tr>
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

[VK_EXT_sample_locations](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_sample_locations.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkSampleLocationsInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSampleLocationsInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetSampleLocationsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
