# vkCmdPushDescriptorSetWithTemplate2KHR(3) Manual Page

## Name

vkCmdPushDescriptorSetWithTemplate2KHR - Pushes descriptor updates into
a command buffer using a descriptor update template



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to use a descriptor update template to specify the push
descriptors to update, call:

``` c
// Provided by VK_KHR_maintenance6 with VK_KHR_push_descriptor
void vkCmdPushDescriptorSetWithTemplate2KHR(
    VkCommandBuffer                             commandBuffer,
    const VkPushDescriptorSetWithTemplateInfoKHR* pPushDescriptorSetWithTemplateInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the descriptors will be
  recorded in.

- `pPushDescriptorSetWithTemplateInfo` is a pointer to a
  `VkPushDescriptorSetWithTemplateInfoKHR` structure.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-parameter"
  id="VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-parameter"></a>
  VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplate2KHR-pPushDescriptorSetWithTemplateInfo-parameter"
  id="VUID-vkCmdPushDescriptorSetWithTemplate2KHR-pPushDescriptorSetWithTemplateInfo-parameter"></a>
  VUID-vkCmdPushDescriptorSetWithTemplate2KHR-pPushDescriptorSetWithTemplateInfo-parameter  
  `pPushDescriptorSetWithTemplateInfo` **must** be a valid pointer to a
  valid
  [VkPushDescriptorSetWithTemplateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetWithTemplateInfoKHR.html)
  structure

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-recording"
  id="VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-recording"></a>
  VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-cmdpool"
  id="VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdPushDescriptorSetWithTemplate2KHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdPushDescriptorSetWithTemplate2KHR-videocoding"
  id="VUID-vkCmdPushDescriptorSetWithTemplate2KHR-videocoding"></a>
  VUID-vkCmdPushDescriptorSetWithTemplate2KHR-videocoding  
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
<td class="tableblock halign-left valign-top"><p>Graphics<br />
Compute</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance6](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance6.html),
[VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPushDescriptorSetWithTemplateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPushDescriptorSetWithTemplateInfoKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPushDescriptorSetWithTemplate2KHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
