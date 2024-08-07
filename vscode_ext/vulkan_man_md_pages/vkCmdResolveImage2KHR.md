# vkCmdResolveImage2(3) Manual Page

## Name

vkCmdResolveImage2 - Resolve regions of an image



## <a href="#_c_specification" class="anchor"></a>C Specification

To resolve a multisample image to a non-multisample image, call:

``` c
// Provided by VK_VERSION_1_3
void vkCmdResolveImage2(
    VkCommandBuffer                             commandBuffer,
    const VkResolveImageInfo2*                  pResolveImageInfo);
```

or the equivalent command

``` c
// Provided by VK_KHR_copy_commands2
void vkCmdResolveImage2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkResolveImageInfo2*                  pResolveImageInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command will be
  recorded.

- `pResolveImageInfo` is a pointer to a
  [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveImageInfo2.html) structure describing
  the resolve parameters.

## <a href="#_description" class="anchor"></a>Description

This command is functionally identical to
[vkCmdResolveImage](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdResolveImage.html), but includes extensible
sub-structures that include `sType` and `pNext` parameters, allowing
them to be more easily extended.

Valid Usage

- <a href="#VUID-vkCmdResolveImage2-commandBuffer-01837"
  id="VUID-vkCmdResolveImage2-commandBuffer-01837"></a>
  VUID-vkCmdResolveImage2-commandBuffer-01837  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `srcImage` **must** not be a protected image

- <a href="#VUID-vkCmdResolveImage2-commandBuffer-01838"
  id="VUID-vkCmdResolveImage2-commandBuffer-01838"></a>
  VUID-vkCmdResolveImage2-commandBuffer-01838  
  If `commandBuffer` is an unprotected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be a protected image

- <a href="#VUID-vkCmdResolveImage2-commandBuffer-01839"
  id="VUID-vkCmdResolveImage2-commandBuffer-01839"></a>
  VUID-vkCmdResolveImage2-commandBuffer-01839  
  If `commandBuffer` is a protected command buffer and
  [`protectedNoFault`](#limits-protectedNoFault) is not supported,
  `dstImage` **must** not be an unprotected image

Valid Usage (Implicit)

- <a href="#VUID-vkCmdResolveImage2-commandBuffer-parameter"
  id="VUID-vkCmdResolveImage2-commandBuffer-parameter"></a>
  VUID-vkCmdResolveImage2-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdResolveImage2-pResolveImageInfo-parameter"
  id="VUID-vkCmdResolveImage2-pResolveImageInfo-parameter"></a>
  VUID-vkCmdResolveImage2-pResolveImageInfo-parameter  
  `pResolveImageInfo` **must** be a valid pointer to a valid
  [VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveImageInfo2.html) structure

- <a href="#VUID-vkCmdResolveImage2-commandBuffer-recording"
  id="VUID-vkCmdResolveImage2-commandBuffer-recording"></a>
  VUID-vkCmdResolveImage2-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdResolveImage2-commandBuffer-cmdpool"
  id="VUID-vkCmdResolveImage2-commandBuffer-cmdpool"></a>
  VUID-vkCmdResolveImage2-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdResolveImage2-renderpass"
  id="VUID-vkCmdResolveImage2-renderpass"></a>
  VUID-vkCmdResolveImage2-renderpass  
  This command **must** only be called outside of a render pass instance

- <a href="#VUID-vkCmdResolveImage2-videocoding"
  id="VUID-vkCmdResolveImage2-videocoding"></a>
  VUID-vkCmdResolveImage2-videocoding  
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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkResolveImageInfo2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkResolveImageInfo2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdResolveImage2"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
