# vkCmdSetViewportShadingRatePaletteNV(3) Manual Page

## Name

vkCmdSetViewportShadingRatePaletteNV - Set shading rate image palettes
dynamically for a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-dynamic-state"
target="_blank" rel="noopener">dynamically set</a> the per-viewport
shading rate image palettes, call:

``` c
// Provided by VK_NV_shading_rate_image
void vkCmdSetViewportShadingRatePaletteNV(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstViewport,
    uint32_t                                    viewportCount,
    const VkShadingRatePaletteNV*               pShadingRatePalettes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `firstViewport` is the index of the first viewport whose shading rate
  palette is updated by the command.

- `viewportCount` is the number of viewports whose shading rate palettes
  are updated by the command.

- `pShadingRatePalettes` is a pointer to an array of
  [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteNV.html) structures
  defining the palette for each viewport.

## <a href="#_description" class="anchor"></a>Description

This command sets the per-viewport shading rate image palettes for
subsequent drawing commands when drawing using <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-objects"
target="_blank" rel="noopener">shader objects</a>, or when the graphics
pipeline is created with
`VK_DYNAMIC_STATE_VIEWPORT_SHADING_RATE_PALETTE_NV` set in
[VkPipelineDynamicStateCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineDynamicStateCreateInfo.html)::`pDynamicStates`.
Otherwise, this state is specified by the
[VkPipelineViewportShadingRateImageStateCreateInfoNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineViewportShadingRateImageStateCreateInfoNV.html)::`pShadingRatePalettes`
values used to create the currently active pipeline.

Valid Usage

- <a href="#VUID-vkCmdSetViewportShadingRatePaletteNV-None-02064"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-None-02064"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-None-02064  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shadingRateImage"
  target="_blank" rel="noopener"><code>shadingRateImage</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02067"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02067"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02067  
  The sum of `firstViewport` and `viewportCount` **must** be between `1`
  and `VkPhysicalDeviceLimits`::`maxViewports`, inclusive

- <a href="#VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02068"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02068"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-firstViewport-02068  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `firstViewport` **must** be `0`

- <a href="#VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-02069"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-02069"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-02069  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-multiViewport"
  target="_blank" rel="noopener"><code>multiViewport</code></a> feature
  is not enabled, `viewportCount` **must** be `1`

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-parameter"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-parameter"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetViewportShadingRatePaletteNV-pShadingRatePalettes-parameter"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-pShadingRatePalettes-parameter"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-pShadingRatePalettes-parameter  
  `pShadingRatePalettes` **must** be a valid pointer to an array of
  `viewportCount` valid
  [VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteNV.html) structures

- <a
  href="#VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-recording"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-recording"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-cmdpool"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdSetViewportShadingRatePaletteNV-videocoding"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-videocoding"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-arraylength"
  id="VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-arraylength"></a>
  VUID-vkCmdSetViewportShadingRatePaletteNV-viewportCount-arraylength  
  `viewportCount` **must** be greater than `0`

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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkShadingRatePaletteNV](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkShadingRatePaletteNV.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetViewportShadingRatePaletteNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
