# vkCmdBindShadingRateImageNV(3) Manual Page

## Name

vkCmdBindShadingRateImageNV - Bind a shading rate image on a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

When shading rate image usage is enabled in the bound pipeline, the
pipeline uses a shading rate image specified by the command:

``` c
// Provided by VK_NV_shading_rate_image
void vkCmdBindShadingRateImageNV(
    VkCommandBuffer                             commandBuffer,
    VkImageView                                 imageView,
    VkImageLayout                               imageLayout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `imageView` is an image view handle specifying the shading rate image.
  `imageView` **may** be set to [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  which is equivalent to specifying a view of an image filled with zero
  values.

- `imageLayout` is the layout that the image subresources accessible
  from `imageView` will be in when the shading rate image is accessed.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdBindShadingRateImageNV-None-02058"
  id="VUID-vkCmdBindShadingRateImageNV-None-02058"></a>
  VUID-vkCmdBindShadingRateImageNV-None-02058  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-shadingRateImage"
  target="_blank" rel="noopener"><code>shadingRateImage</code></a>
  feature **must** be enabled

- <a href="#VUID-vkCmdBindShadingRateImageNV-imageView-02059"
  id="VUID-vkCmdBindShadingRateImageNV-imageView-02059"></a>
  VUID-vkCmdBindShadingRateImageNV-imageView-02059  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle of type
  `VK_IMAGE_VIEW_TYPE_2D` or `VK_IMAGE_VIEW_TYPE_2D_ARRAY`

- <a href="#VUID-vkCmdBindShadingRateImageNV-imageView-02060"
  id="VUID-vkCmdBindShadingRateImageNV-imageView-02060"></a>
  VUID-vkCmdBindShadingRateImageNV-imageView-02060  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** have a format of `VK_FORMAT_R8_UINT`

- <a href="#VUID-vkCmdBindShadingRateImageNV-imageView-02061"
  id="VUID-vkCmdBindShadingRateImageNV-imageView-02061"></a>
  VUID-vkCmdBindShadingRateImageNV-imageView-02061  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** have been created with a `usage` value including
  `VK_IMAGE_USAGE_SHADING_RATE_IMAGE_BIT_NV`

- <a href="#VUID-vkCmdBindShadingRateImageNV-imageView-02062"
  id="VUID-vkCmdBindShadingRateImageNV-imageView-02062"></a>
  VUID-vkCmdBindShadingRateImageNV-imageView-02062  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** match the actual
  [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html) of each subresource accessible
  from `imageView` at the time the subresource is accessed

- <a href="#VUID-vkCmdBindShadingRateImageNV-imageLayout-02063"
  id="VUID-vkCmdBindShadingRateImageNV-imageLayout-02063"></a>
  VUID-vkCmdBindShadingRateImageNV-imageLayout-02063  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** be `VK_IMAGE_LAYOUT_SHADING_RATE_OPTIMAL_NV` or
  `VK_IMAGE_LAYOUT_GENERAL`

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindShadingRateImageNV-commandBuffer-parameter"
  id="VUID-vkCmdBindShadingRateImageNV-commandBuffer-parameter"></a>
  VUID-vkCmdBindShadingRateImageNV-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindShadingRateImageNV-imageView-parameter"
  id="VUID-vkCmdBindShadingRateImageNV-imageView-parameter"></a>
  VUID-vkCmdBindShadingRateImageNV-imageView-parameter  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-vkCmdBindShadingRateImageNV-imageLayout-parameter"
  id="VUID-vkCmdBindShadingRateImageNV-imageLayout-parameter"></a>
  VUID-vkCmdBindShadingRateImageNV-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-vkCmdBindShadingRateImageNV-commandBuffer-recording"
  id="VUID-vkCmdBindShadingRateImageNV-commandBuffer-recording"></a>
  VUID-vkCmdBindShadingRateImageNV-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindShadingRateImageNV-commandBuffer-cmdpool"
  id="VUID-vkCmdBindShadingRateImageNV-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindShadingRateImageNV-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBindShadingRateImageNV-videocoding"
  id="VUID-vkCmdBindShadingRateImageNV-videocoding"></a>
  VUID-vkCmdBindShadingRateImageNV-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindShadingRateImageNV-commonparent"
  id="VUID-vkCmdBindShadingRateImageNV-commonparent"></a>
  VUID-vkCmdBindShadingRateImageNV-commonparent  
  Both of `commandBuffer`, and `imageView` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindShadingRateImageNV"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
