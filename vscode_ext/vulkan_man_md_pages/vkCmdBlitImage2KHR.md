# vkCmdBlitImage2(3) Manual Page

## Name

vkCmdBlitImage2 - Copy regions of an image, potentially performing
format conversion,



## <a href="#_c_specification" class="anchor"></a>C Specification

To copy regions of a source image into a destination image, potentially
performing format conversion, arbitrary scaling, and filtering, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdBlitImage2(
    VkCommandBuffer                             commandBuffer,
    const VkBlitImageInfo2*                     pBlitImageInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_copy_commands2
void vkCmdBlitImage2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkBlitImageInfo2*                     pBlitImageInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pBlitImageInfo` is a pointer to a
  [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html) structure describing the
  blit parameters.

## <a href="#_description" class="anchor"></a>Description

This command is functionally identical to
[vkCmdBlitImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBlitImage.html), but includes extensible
sub-structures that include `sType` and `pNext` parameters, allowing
them to be more easily extended.

Valid Usage

- <a href="#VUID-vkCmdBlitImage2-commandBuffer-01834"
  id="VUID-vkCmdBlitImage2-commandBuffer-01834"></a>
  VUID-vkCmdBlitImage2-commandBuffer-01834  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcImage` **must** not be a protected image

- <a href="#VUID-vkCmdBlitImage2-commandBuffer-01835"
  id="VUID-vkCmdBlitImage2-commandBuffer-01835"></a>
  VUID-vkCmdBlitImage2-commandBuffer-01835  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdBlitImage2-commandBuffer-01836"
  id="VUID-vkCmdBlitImage2-commandBuffer-01836"></a>
  VUID-vkCmdBlitImage2-commandBuffer-01836  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be an unprotected image

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBlitImage2-commandBuffer-parameter"
  id="VUID-vkCmdBlitImage2-commandBuffer-parameter"></a>
  VUID-vkCmdBlitImage2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBlitImage2-pBlitImageInfo-parameter"
  id="VUID-vkCmdBlitImage2-pBlitImageInfo-parameter"></a>
  VUID-vkCmdBlitImage2-pBlitImageInfo-parameter  
  `pBlitImageInfo` **must** be a valid pointer to a valid
  [VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html) structure

- <a href="#VUID-vkCmdBlitImage2-commandBuffer-recording"
  id="VUID-vkCmdBlitImage2-commandBuffer-recording"></a>
  VUID-vkCmdBlitImage2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBlitImage2-commandBuffer-cmdpool"
  id="VUID-vkCmdBlitImage2-commandBuffer-cmdpool"></a>
  VUID-vkCmdBlitImage2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBlitImage2-renderpass"
  id="VUID-vkCmdBlitImage2-renderpass"></a>
  VUID-vkCmdBlitImage2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdBlitImage2-videocoding"
  id="VUID-vkCmdBlitImage2-videocoding"></a>
  VUID-vkCmdBlitImage2-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Outside</p></td>
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>Action</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_copy_commands2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_copy_commands2.html),
[VK_VERSION_1_3](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_3.html),
[VkBlitImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBlitImageInfo2.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBlitImage2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
