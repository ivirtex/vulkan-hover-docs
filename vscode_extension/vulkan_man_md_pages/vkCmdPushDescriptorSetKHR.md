# vkCmdPushDescriptorSetKHR(3) Manual Page

## Name

vkCmdPushDescriptorSetKHR - Pushes descriptor updates into a command
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

In addition to allocating descriptor sets and binding them to a command
buffer, an application **can** record descriptor updates into the
command buffer.

To push descriptor updates into a command buffer, call:

``` c
// Provided by VK_KHR_push_descriptor
void vkCmdPushDescriptorSetKHR(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    set,
    uint32_t                                    descriptorWriteCount,
    const VkWriteDescriptorSet*                 pDescriptorWrites);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the descriptors will be
  recorded in.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) indicating the type of
  the pipeline that will use the descriptors. There is a separate set of
  push descriptor bindings for each pipeline type, so binding one does
  not disturb the others.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings.

- `set` is the set number of the descriptor set in the pipeline layout
  that will be updated.

- `descriptorWriteCount` is the number of elements in the
  `pDescriptorWrites` array.

- `pDescriptorWrites` is a pointer to an array of
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) structures
  describing the descriptors to be updated.

## <a href="#_description" class="anchor"></a>Description

*Push descriptors* are a small bank of descriptors whose storage is
internally managed by the command buffer rather than being written into
a descriptor set and later bound to a command buffer. Push descriptors
allow for incremental updates of descriptors without managing the
lifetime of descriptor sets.

When a command buffer begins recording, all push descriptors are
undefined. Push descriptors **can** be updated incrementally and cause
shaders to use the updated descriptors for subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-bindpoint-commands"
target="_blank" rel="noopener">bound pipeline commands</a> with the
pipeline type set by `pipelineBindPoint` until the descriptor is
overwritten, or else until the set is disturbed as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">Pipeline Layout Compatibility</a>. When
the set is disturbed or push descriptors with a different descriptor set
layout are set, all push descriptors are undefined.

Push descriptors that are <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#shaders-staticuse"
target="_blank" rel="noopener">statically used</a> by a pipeline
**must** not be undefined at the time that a drawing or dispatching
command is recorded to execute using that pipeline. This includes
immutable sampler descriptors, which **must** be pushed before they are
accessed by a pipeline (the immutable samplers are pushed, rather than
the samplers in `pDescriptorWrites`). Push descriptors that are not
statically used **can** remain undefined.

Push descriptors do not use dynamic offsets. Instead, the corresponding
non-dynamic descriptor types **can** be used and the `offset` member of
[VkDescriptorBufferInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBufferInfo.html) **can** be changed
each time the descriptor is written.

Each element of `pDescriptorWrites` is interpreted as in
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html), except the `dstSet`
member is ignored.

To push an immutable sampler, use a
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) with `dstBinding` and
`dstArrayElement` selecting the immutable sampler’s binding. If the
descriptor type is `VK_DESCRIPTOR_TYPE_SAMPLER`, the `pImageInfo`
parameter is ignored and the immutable sampler is taken from the push
descriptor set layout in the pipeline layout. If the descriptor type is
`VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`, the `sampler` member of the
`pImageInfo` parameter is ignored and the immutable sampler is taken
from the push descriptor set layout in the pipeline layout.

Valid Usage

- <a href="#VUID-vkCmdPushDescriptorSetKHR-set-00364"
  id="VUID-vkCmdPushDescriptorSetKHR-set-00364"></a>
  VUID-vkCmdPushDescriptorSetKHR-set-00364  
  `set` **must** be less than
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-vkCmdPushDescriptorSetKHR-set-00365"
  id="VUID-vkCmdPushDescriptorSetKHR-set-00365"></a>
  VUID-vkCmdPushDescriptorSetKHR-set-00365  
  `set` **must** be the unique set number in the pipeline layout that
  uses a descriptor set layout that was created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR`

- <a href="#VUID-vkCmdPushDescriptorSetKHR-pDescriptorWrites-06494"
  id="VUID-vkCmdPushDescriptorSetKHR-pDescriptorWrites-06494"></a>
  VUID-vkCmdPushDescriptorSetKHR-pDescriptorWrites-06494  
  For each element i where `pDescriptorWrites`\[i\].`descriptorType` is
  `VK_DESCRIPTOR_TYPE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER`,
  `VK_DESCRIPTOR_TYPE_SAMPLED_IMAGE`,
  `VK_DESCRIPTOR_TYPE_STORAGE_IMAGE`, or
  `VK_DESCRIPTOR_TYPE_INPUT_ATTACHMENT`,
  `pDescriptorWrites`\[i\].`pImageInfo` **must** be a valid pointer to
  an array of `pDescriptorWrites`\[i\].`descriptorCount` valid
  `VkDescriptorImageInfo` structures

- <a href="#VUID-vkCmdPushDescriptorSetKHR-pipelineBindPoint-00363"
  id="VUID-vkCmdPushDescriptorSetKHR-pipelineBindPoint-00363"></a>
  VUID-vkCmdPushDescriptorSetKHR-pipelineBindPoint-00363  
  `pipelineBindPoint` **must** be supported by the `commandBuffer`’s
  parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- <a href="#VUID-vkCmdPushDescriptorSetKHR-commandBuffer-parameter"
  id="VUID-vkCmdPushDescriptorSetKHR-commandBuffer-parameter"></a>
  VUID-vkCmdPushDescriptorSetKHR-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdPushDescriptorSetKHR-pipelineBindPoint-parameter"
  id="VUID-vkCmdPushDescriptorSetKHR-pipelineBindPoint-parameter"></a>
  VUID-vkCmdPushDescriptorSetKHR-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-vkCmdPushDescriptorSetKHR-layout-parameter"
  id="VUID-vkCmdPushDescriptorSetKHR-layout-parameter"></a>
  VUID-vkCmdPushDescriptorSetKHR-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-vkCmdPushDescriptorSetKHR-pDescriptorWrites-parameter"
  id="VUID-vkCmdPushDescriptorSetKHR-pDescriptorWrites-parameter"></a>
  VUID-vkCmdPushDescriptorSetKHR-pDescriptorWrites-parameter  
  `pDescriptorWrites` **must** be a valid pointer to an array of
  `descriptorWriteCount` valid
  [VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html) structures

- <a href="#VUID-vkCmdPushDescriptorSetKHR-commandBuffer-recording"
  id="VUID-vkCmdPushDescriptorSetKHR-commandBuffer-recording"></a>
  VUID-vkCmdPushDescriptorSetKHR-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdPushDescriptorSetKHR-commandBuffer-cmdpool"
  id="VUID-vkCmdPushDescriptorSetKHR-commandBuffer-cmdpool"></a>
  VUID-vkCmdPushDescriptorSetKHR-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdPushDescriptorSetKHR-videocoding"
  id="VUID-vkCmdPushDescriptorSetKHR-videocoding"></a>
  VUID-vkCmdPushDescriptorSetKHR-videocoding  
  This command **must** only be called outside of a video coding scope

- <a
  href="#VUID-vkCmdPushDescriptorSetKHR-descriptorWriteCount-arraylength"
  id="VUID-vkCmdPushDescriptorSetKHR-descriptorWriteCount-arraylength"></a>
  VUID-vkCmdPushDescriptorSetKHR-descriptorWriteCount-arraylength  
  `descriptorWriteCount` **must** be greater than `0`

- <a href="#VUID-vkCmdPushDescriptorSetKHR-commonparent"
  id="VUID-vkCmdPushDescriptorSetKHR-commonparent"></a>
  VUID-vkCmdPushDescriptorSetKHR-commonparent  
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

[VK_KHR_push_descriptor](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_push_descriptor.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html),
[VkWriteDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkWriteDescriptorSet.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdPushDescriptorSetKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
