# vkCmdBindDescriptorSets(3) Manual Page

## Name

vkCmdBindDescriptorSets - Binds descriptor sets to a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To bind one or more descriptor sets to a command buffer, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdBindDescriptorSets(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    firstSet,
    uint32_t                                    descriptorSetCount,
    const VkDescriptorSet*                      pDescriptorSets,
    uint32_t                                    dynamicOffsetCount,
    const uint32_t*                             pDynamicOffsets);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer that the descriptor sets will be
  bound to.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) indicating the type of
  the pipeline that will use the descriptors. There is a separate set of
  bind points for each pipeline type, so binding one does not disturb
  the others.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings.

- `firstSet` is the set number of the first descriptor set to be bound.

- `descriptorSetCount` is the number of elements in the
  `pDescriptorSets` array.

- `pDescriptorSets` is a pointer to an array of handles to
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) objects describing the
  descriptor sets to bind to.

- `dynamicOffsetCount` is the number of dynamic offsets in the
  `pDynamicOffsets` array.

- `pDynamicOffsets` is a pointer to an array of `uint32_t` values
  specifying dynamic offsets.

## <a href="#_description" class="anchor"></a>Description

`vkCmdBindDescriptorSets` binds descriptor sets
`pDescriptorSets`\[0..`descriptorSetCount`-1\] to set numbers
\[`firstSet`..`firstSet`+`descriptorSetCount`-1\] for subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-bindpoint-commands"
target="_blank" rel="noopener">bound pipeline commands</a> set by
`pipelineBindPoint`. Any bindings that were previously applied via these
sets , or calls to
[vkCmdSetDescriptorBufferOffsetsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdSetDescriptorBufferOffsetsEXT.html)
or
[vkCmdBindDescriptorBufferEmbeddedSamplersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBufferEmbeddedSamplersEXT.html),
are no longer valid.

Once bound, a descriptor set affects rendering of subsequent commands
that interact with the given pipeline type in the command buffer until
either a different set is bound to the same set number, or the set is
disturbed as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">Pipeline Layout Compatibility</a>.

A compatible descriptor set **must** be bound for all set numbers that
any shaders in a pipeline access, at the time that a drawing or
dispatching command is recorded to execute using that pipeline. However,
if none of the shaders in a pipeline statically use any bindings with a
particular set number, then no descriptor set need be bound for that set
number, even if the pipeline layout includes a non-trivial descriptor
set layout for that set number.

When consuming a descriptor, a descriptor is considered valid if the
descriptor is not undefined as described by <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptor-set-initial-state"
target="_blank" rel="noopener">descriptor set allocation</a>. If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
is enabled, a null descriptor is also considered valid. A descriptor
that was disturbed by <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">Pipeline Layout Compatibility</a>, or was
never bound by `vkCmdBindDescriptorSets` is not considered valid. If a
pipeline accesses a descriptor either statically or dynamically
depending on the
[VkDescriptorBindingFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorBindingFlagBits.html), the
consuming descriptor type in the pipeline **must** match the
[VkDescriptorType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorType.html) in
[VkDescriptorSetLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayoutCreateInfo.html)
for the descriptor to be considered valid. If a descriptor is a mutable
descriptor, the consuming descriptor type in the pipeline **must** match
the active descriptor type for the descriptor to be considered valid.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>Further validation may be carried out beyond validation for
descriptor types, e.g. <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#textures-input-validation"
target="_blank" rel="noopener">Texel Input Validation</a>.</p></td>
</tr>
</tbody>
</table>

If any of the sets being bound include dynamic uniform or storage
buffers, then `pDynamicOffsets` includes one element for each array
element in each dynamic descriptor type binding in each set. Values are
taken from `pDynamicOffsets` in an order such that all entries for set N
come before set N+1; within a set, entries are ordered by the binding
numbers in the descriptor set layouts; and within a binding array,
elements are in order. `dynamicOffsetCount` **must** equal the total
number of dynamic descriptors in the sets being bound.

The effective offset used for dynamic uniform and storage buffer
bindings is the sum of the relative offset taken from `pDynamicOffsets`,
and the base address of the buffer plus base offset in the descriptor
set. The range of the dynamic uniform and storage buffer bindings is the
buffer range as specified in the descriptor set.

Each of the `pDescriptorSets` **must** be compatible with the pipeline
layout specified by `layout`. The layout used to program the bindings
**must** also be compatible with the pipeline used in subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-bindpoint-commands"
target="_blank" rel="noopener">bound pipeline commands</a> with that
pipeline type, as defined in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">Pipeline Layout Compatibility</a>
section.

The descriptor set contents bound by a call to `vkCmdBindDescriptorSets`
**may** be consumed at the following times:

- For descriptor bindings created with the
  `VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT` bit set, the contents
  **may** be consumed when the command buffer is submitted to a queue,
  or during shader execution of the resulting draws and dispatches, or
  any time in between. Otherwise,

- during host execution of the command, or during shader execution of
  the resulting draws and dispatches, or any time in between.

Thus, the contents of a descriptor set binding **must** not be altered
(overwritten by an update command, or freed) between the first point in
time that it **may** be consumed, and when the command completes
executing on the queue.

The contents of `pDynamicOffsets` are consumed immediately during
execution of `vkCmdBindDescriptorSets`. Once all pending uses have
completed, it is legal to update and reuse a descriptor set.

Valid Usage

- <a href="#VUID-vkCmdBindDescriptorSets-pDescriptorSets-00358"
  id="VUID-vkCmdBindDescriptorSets-pDescriptorSets-00358"></a>
  VUID-vkCmdBindDescriptorSets-pDescriptorSets-00358  
  Each element of `pDescriptorSets` **must** have been allocated with a
  `VkDescriptorSetLayout` that matches (is the same as, or identically
  defined as) the `VkDescriptorSetLayout` at set *n* in `layout`, where
  *n* is the sum of `firstSet` and the index into `pDescriptorSets`

- <a href="#VUID-vkCmdBindDescriptorSets-dynamicOffsetCount-00359"
  id="VUID-vkCmdBindDescriptorSets-dynamicOffsetCount-00359"></a>
  VUID-vkCmdBindDescriptorSets-dynamicOffsetCount-00359  
  `dynamicOffsetCount` **must** be equal to the total number of dynamic
  descriptors in `pDescriptorSets`

- <a href="#VUID-vkCmdBindDescriptorSets-firstSet-00360"
  id="VUID-vkCmdBindDescriptorSets-firstSet-00360"></a>
  VUID-vkCmdBindDescriptorSets-firstSet-00360  
  The sum of `firstSet` and `descriptorSetCount` **must** be less than
  or equal to
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01971"
  id="VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01971"></a>
  VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01971  
  Each element of `pDynamicOffsets` which corresponds to a descriptor
  binding with type `VK_DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC` **must**
  be a multiple of
  `VkPhysicalDeviceLimits`::`minUniformBufferOffsetAlignment`

- <a href="#VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01972"
  id="VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01972"></a>
  VUID-vkCmdBindDescriptorSets-pDynamicOffsets-01972  
  Each element of `pDynamicOffsets` which corresponds to a descriptor
  binding with type `VK_DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC` **must**
  be a multiple of
  `VkPhysicalDeviceLimits`::`minStorageBufferOffsetAlignment`

- <a href="#VUID-vkCmdBindDescriptorSets-pDescriptorSets-01979"
  id="VUID-vkCmdBindDescriptorSets-pDescriptorSets-01979"></a>
  VUID-vkCmdBindDescriptorSets-pDescriptorSets-01979  
  For each dynamic uniform or storage buffer binding in
  `pDescriptorSets`, the sum of the [effective
  offset](#dynamic-effective-offset) and the range of the binding
  **must** be less than or equal to the size of the buffer

- <a href="#VUID-vkCmdBindDescriptorSets-pDescriptorSets-06715"
  id="VUID-vkCmdBindDescriptorSets-pDescriptorSets-06715"></a>
  VUID-vkCmdBindDescriptorSets-pDescriptorSets-06715  
  For each dynamic uniform or storage buffer binding in
  `pDescriptorSets`, if the range was set with `VK_WHOLE_SIZE` then
  `pDynamicOffsets` which corresponds to the descriptor binding **must**
  be 0

- <a href="#VUID-vkCmdBindDescriptorSets-pDescriptorSets-04616"
  id="VUID-vkCmdBindDescriptorSets-pDescriptorSets-04616"></a>
  VUID-vkCmdBindDescriptorSets-pDescriptorSets-04616  
  Each element of `pDescriptorSets` **must** not have been allocated
  from a `VkDescriptorPool` with the
  `VK_DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT` flag set

- <a href="#VUID-vkCmdBindDescriptorSets-pDescriptorSets-06563"
  id="VUID-vkCmdBindDescriptorSets-pDescriptorSets-06563"></a>
  VUID-vkCmdBindDescriptorSets-pDescriptorSets-06563  
  If [`graphicsPipelineLibrary`](#features-graphicsPipelineLibrary) is
  not enabled, each element of `pDescriptorSets` **must** be a valid
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html)

- <a href="#VUID-vkCmdBindDescriptorSets-pDescriptorSets-08010"
  id="VUID-vkCmdBindDescriptorSets-pDescriptorSets-08010"></a>
  VUID-vkCmdBindDescriptorSets-pDescriptorSets-08010  
  Each element of `pDescriptorSets` **must** have been allocated with a
  `VkDescriptorSetLayout` which was not created with
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT`

- <a href="#VUID-vkCmdBindDescriptorSets-pipelineBindPoint-00361"
  id="VUID-vkCmdBindDescriptorSets-pipelineBindPoint-00361"></a>
  VUID-vkCmdBindDescriptorSets-pipelineBindPoint-00361  
  `pipelineBindPoint` **must** be supported by the `commandBuffer`’s
  parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindDescriptorSets-commandBuffer-parameter"
  id="VUID-vkCmdBindDescriptorSets-commandBuffer-parameter"></a>
  VUID-vkCmdBindDescriptorSets-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindDescriptorSets-pipelineBindPoint-parameter"
  id="VUID-vkCmdBindDescriptorSets-pipelineBindPoint-parameter"></a>
  VUID-vkCmdBindDescriptorSets-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-vkCmdBindDescriptorSets-layout-parameter"
  id="VUID-vkCmdBindDescriptorSets-layout-parameter"></a>
  VUID-vkCmdBindDescriptorSets-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a href="#VUID-vkCmdBindDescriptorSets-pDescriptorSets-parameter"
  id="VUID-vkCmdBindDescriptorSets-pDescriptorSets-parameter"></a>
  VUID-vkCmdBindDescriptorSets-pDescriptorSets-parameter  
  `pDescriptorSets` **must** be a valid pointer to an array of
  `descriptorSetCount` valid or [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  [VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html) handles

- <a href="#VUID-vkCmdBindDescriptorSets-pDynamicOffsets-parameter"
  id="VUID-vkCmdBindDescriptorSets-pDynamicOffsets-parameter"></a>
  VUID-vkCmdBindDescriptorSets-pDynamicOffsets-parameter  
  If `dynamicOffsetCount` is not `0`, `pDynamicOffsets` **must** be a
  valid pointer to an array of `dynamicOffsetCount` `uint32_t` values

- <a href="#VUID-vkCmdBindDescriptorSets-commandBuffer-recording"
  id="VUID-vkCmdBindDescriptorSets-commandBuffer-recording"></a>
  VUID-vkCmdBindDescriptorSets-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindDescriptorSets-commandBuffer-cmdpool"
  id="VUID-vkCmdBindDescriptorSets-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindDescriptorSets-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdBindDescriptorSets-videocoding"
  id="VUID-vkCmdBindDescriptorSets-videocoding"></a>
  VUID-vkCmdBindDescriptorSets-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindDescriptorSets-descriptorSetCount-arraylength"
  id="VUID-vkCmdBindDescriptorSets-descriptorSetCount-arraylength"></a>
  VUID-vkCmdBindDescriptorSets-descriptorSetCount-arraylength  
  `descriptorSetCount` **must** be greater than `0`

- <a href="#VUID-vkCmdBindDescriptorSets-commonparent"
  id="VUID-vkCmdBindDescriptorSets-commonparent"></a>
  VUID-vkCmdBindDescriptorSets-commonparent  
  Each of `commandBuffer`, `layout`, and the elements of
  `pDescriptorSets` that are valid handles of non-ignored parameters
  **must** have been created, allocated, or retrieved from the same
  [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDescriptorSet](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSet.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindDescriptorSets"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
