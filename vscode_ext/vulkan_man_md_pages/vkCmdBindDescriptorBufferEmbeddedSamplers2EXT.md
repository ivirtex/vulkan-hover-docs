# vkCmdBindDescriptorBufferEmbeddedSamplers2EXT(3) Manual Page

## Name

vkCmdBindDescriptorBufferEmbeddedSamplers2EXT - Setting embedded
immutable samplers offsets in a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to bind an embedded immutable sampler set to a command
buffer, call:

``` c
// Provided by VK_KHR_maintenance6 with VK_EXT_descriptor_buffer
void vkCmdBindDescriptorBufferEmbeddedSamplers2EXT(
    VkCommandBuffer                             commandBuffer,
    const VkBindDescriptorBufferEmbeddedSamplersInfoEXT* pBindDescriptorBufferEmbeddedSamplersInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the embedded immutable
  samplers will be bound to.

- `pBindDescriptorBufferEmbeddedSamplersInfo` is a pointer to a
  `VkBindDescriptorBufferEmbeddedSamplersInfoEXT` structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-descriptorBuffer-09472"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-descriptorBuffer-09472"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-descriptorBuffer-09472  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-09473"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-09473"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-09473  
  Each bit in `pBindDescriptorBufferEmbeddedSamplersInfo->stageFlags`
  **must** be a stage supported by the `commandBuffer`’s parent
  `VkCommandPool`’s queue family

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-parameter"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-parameter"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-parameter"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-parameter"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-pBindDescriptorBufferEmbeddedSamplersInfo-parameter  
  `pBindDescriptorBufferEmbeddedSamplersInfo` **must** be a valid
  pointer to a valid
  [VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html)
  structure

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-recording"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-recording"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-videocoding"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-videocoding"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplers2EXT-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VkBindDescriptorBufferEmbeddedSamplersInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBindDescriptorBufferEmbeddedSamplersInfoEXT.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindDescriptorBufferEmbeddedSamplers2EXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
