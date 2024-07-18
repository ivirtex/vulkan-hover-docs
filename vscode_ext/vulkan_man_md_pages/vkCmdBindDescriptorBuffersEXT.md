# vkCmdBindDescriptorBuffersEXT(3) Manual Page

## Name

vkCmdBindDescriptorBuffersEXT - Binding descriptor buffers to a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To bind descriptor buffers to a command buffer, call:

``` c
// Provided by VK_EXT_descriptor_buffer
void vkCmdBindDescriptorBuffersEXT(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    bufferCount,
    const VkDescriptorBufferBindingInfoEXT*     pBindingInfos);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the descriptor buffers will
  be bound to.

- `bufferCount` is the number of elements in the `pBindingInfos` array.

- `pBindingInfos` is a pointer to an array of
  [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferBindingInfoEXT.html)
  structures.

## <a href="#_description" class="anchor"></a>Description

`vkCmdBindDescriptorBuffersEXT` causes any offsets previously set by
[vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html)
that use the bindings numbered \[`0`.. `bufferCount`-1\] to be no longer
valid for subsequent bound pipeline commands. Any previously bound
buffers at binding points greater than or equal to `bufferCount` are
unbound.

Valid Usage

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-None-08047"
  id="VUID-vkCmdBindDescriptorBuffersEXT-None-08047"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-None-08047  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdBindDescriptorBuffersEXT-maxSamplerDescriptorBufferBindings-08048"
  id="VUID-vkCmdBindDescriptorBuffersEXT-maxSamplerDescriptorBufferBindings-08048"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-maxSamplerDescriptorBufferBindings-08048  
  There **must** be no more than
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxSamplerDescriptorBufferBindings`
  descriptor buffers containing sampler descriptor data bound

- <a
  href="#VUID-vkCmdBindDescriptorBuffersEXT-maxResourceDescriptorBufferBindings-08049"
  id="VUID-vkCmdBindDescriptorBuffersEXT-maxResourceDescriptorBufferBindings-08049"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-maxResourceDescriptorBufferBindings-08049  
  There **must** be no more than
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxResourceDescriptorBufferBindings`
  descriptor buffers containing resource descriptor data bound

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-None-08050"
  id="VUID-vkCmdBindDescriptorBuffersEXT-None-08050"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-None-08050  
  There **must** be no more than `1` descriptor buffer bound that was
  created with the
  `VK_BUFFER_USAGE_PUSH_DESCRIPTORS_DESCRIPTOR_BUFFER_BIT_EXT` bit set

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-08051"
  id="VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-08051"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-08051  
  `bufferCount` **must** be less than or equal to
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxDescriptorBufferBindings`

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08052"
  id="VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08052"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08052  
  For any element of `pBindingInfos`, if the buffer from which `address`
  was queried is non-sparse then it **must** be bound completely and
  contiguously to a single [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08053"
  id="VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08053"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08053  
  For any element of `pBindingInfos`, the buffer from which `address`
  was queried **must** have been created with the
  `VK_BUFFER_USAGE_SAMPLER_DESCRIPTOR_BUFFER_BIT_EXT` bit set if it
  contains sampler descriptor data

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08054"
  id="VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08054"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08054  
  For any element of `pBindingInfos`, the buffer from which `address`
  was queried **must** have been created with the
  `VK_BUFFER_USAGE_RESOURCE_DESCRIPTOR_BUFFER_BIT_EXT` bit set if it
  contains resource descriptor data

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08055"
  id="VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08055"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-08055  
  For any element of `pBindingInfos`, `usage` **must** match the buffer
  from which `address` was queried

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-parameter"
  id="VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-parameter"
  id="VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-parameter"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-pBindingInfos-parameter  
  `pBindingInfos` **must** be a valid pointer to an array of
  `bufferCount` valid
  [VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferBindingInfoEXT.html)
  structures

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-recording"
  id="VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-recording"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-videocoding"
  id="VUID-vkCmdBindDescriptorBuffersEXT-videocoding"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-arraylength"
  id="VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-arraylength"></a>
  VUID-vkCmdBindDescriptorBuffersEXT-bufferCount-arraylength  
  `bufferCount` **must** be greater than `0`

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

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDescriptorBufferBindingInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferBindingInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindDescriptorBuffersEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
