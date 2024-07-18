# vkCmdBindInvocationMaskHUAWEI(3) Manual Page

## Name

vkCmdBindInvocationMaskHUAWEI - Bind an invocation mask image on a
command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

When invocation mask image usage is enabled in the bound ray tracing
pipeline, the pipeline uses an invocation mask image specified by the
command:

``` c
// Provided by VK_HUAWEI_invocation_mask
void vkCmdBindInvocationMaskHUAWEI(
    VkCommandBuffer                             commandBuffer,
    VkImageView                                 imageView,
    VkImageLayout                               imageLayout);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded

- `imageView` is an image view handle specifying the invocation mask
  image `imageView` **may** be set to
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), which is equivalent to
  specifying a view of an image filled with ones value.

- `imageLayout` is the layout that the image subresources accessible
  from `imageView` will be in when the invocation mask image is accessed

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-None-04976"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-None-04976"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-None-04976  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-invocationMask"
  target="_blank" rel="noopener"><code>invocationMask</code></a> feature
  **must** be enabled

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04977"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04977"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04977  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle of type
  `VK_IMAGE_VIEW_TYPE_2D`

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04978"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04978"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04978  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** have a format of `VK_FORMAT_R8_UINT`

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04979"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04979"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04979  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), it
  **must** have been created with
  `VK_IMAGE_USAGE_INVOCATION_MASK_BIT_HUAWEI` set

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04980"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04980"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-imageView-04980  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageLayout` **must** be `VK_IMAGE_LAYOUT_GENERAL`

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-width-04981"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-width-04981"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-width-04981  
  Thread mask image resolution **must** match the `width` and `height`
  in [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html)

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-None-04982"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-None-04982"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-None-04982  
  Each element in the invocation mask image **must** have the value `0`
  or `1`. The value 1 means the invocation is active

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-depth-04983"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-depth-04983"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-depth-04983  
  `depth` in [vkCmdTraceRaysKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdTraceRaysKHR.html) **must** be 1

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-parameter"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-parameter"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-imageView-parameter"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-imageView-parameter"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-imageView-parameter  
  If `imageView` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  `imageView` **must** be a valid [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html) handle

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-imageLayout-parameter"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-imageLayout-parameter"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-imageLayout-parameter  
  `imageLayout` **must** be a valid [VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html)
  value

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-recording"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-recording"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-cmdpool"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support compute operations

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-renderpass"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-renderpass"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-videocoding"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-videocoding"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindInvocationMaskHUAWEI-commonparent"
  id="VUID-vkCmdBindInvocationMaskHUAWEI-commonparent"></a>
  VUID-vkCmdBindInvocationMaskHUAWEI-commonparent  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_HUAWEI_invocation_mask](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_HUAWEI_invocation_mask.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkImageLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageLayout.html), [VkImageView](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkImageView.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindInvocationMaskHUAWEI"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
