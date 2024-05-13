# vkCmdBindDescriptorBufferEmbeddedSamplersEXT(3) Manual Page

## Name

vkCmdBindDescriptorBufferEmbeddedSamplersEXT - Setting embedded
immutable samplers offsets in a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To bind an embedded immutable sampler set to a command buffer, call:

``` c
// Provided by VK_EXT_descriptor_buffer
void vkCmdBindDescriptorBufferEmbeddedSamplersEXT(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    set);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the embedded immutable
  samplers will be bound to.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) indicating the type of
  the pipeline that will use the embedded immutable samplers.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings.

- `set` is the number of the set to be bound.

## <a href="#_description" class="anchor"></a>Description

`vkCmdBindDescriptorBufferEmbeddedSamplersEXT` binds the embedded
immutable samplers in `set` of `layout` to `set` for the command buffer
for subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-bindpoint-commands"
target="_blank" rel="noopener">bound pipeline commands</a> set by
`pipelineBindPoint`. Any previous binding to this set by
[vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html)
or this command is overwritten. Any sets that were last bound by a call
to [vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html) are
invalidated upon calling this command. Other sets will also be
invalidated upon calling this command if `layout` differs from the
pipeline layout used to bind those other sets, as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">Pipeline Layout Compatibility</a>.

Valid Usage

- <a href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08070"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08070"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08070  
  The [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) at index `set`
  when `layout` was created **must** have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT`
  bit set

- <a href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08071"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08071"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-set-08071  
  `set` **must** be less than or equal to
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-None-08068"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-None-08068"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-None-08068  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-08069"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-08069"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-08069  
  `pipelineBindPoint` **must** be supported by the `commandBuffer`’s
  parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-parameter"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-parameter"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-parameter"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-parameter"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-layout-parameter"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-layout-parameter"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-recording"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-recording"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-videocoding"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-videocoding"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commonparent"
  id="VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commonparent"></a>
  VUID-vkCmdBindDescriptorBufferEmbeddedSamplersEXT-commonparent  
  Both of `commandBuffer`, and `layout` **must** have been created,
  allocated, or retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindDescriptorBufferEmbeddedSamplersEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
