# vkCmdSetDescriptorBufferOffsetsEXT(3) Manual Page

## Name

vkCmdSetDescriptorBufferOffsetsEXT - Setting descriptor buffer offsets
in a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To set descriptor buffer offsets in a command buffer, call:

``` c
// Provided by VK_EXT_descriptor_buffer
void vkCmdSetDescriptorBufferOffsetsEXT(
    VkCommandBuffer                             commandBuffer,
    VkPipelineBindPoint                         pipelineBindPoint,
    VkPipelineLayout                            layout,
    uint32_t                                    firstSet,
    uint32_t                                    setCount,
    const uint32_t*                             pBufferIndices,
    const VkDeviceSize*                         pOffsets);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer in which the descriptor buffer
  offsets will be set.

- `pipelineBindPoint` is a
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) indicating the type of
  the pipeline that will use the descriptors.

- `layout` is a [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html) object used to
  program the bindings.

- `firstSet` is the number of the first set to be bound.

- `setCount` is the number of elements in the `pBufferIndices` and
  `pOffsets` arrays.

- `pBufferIndices` is a pointer to an array of indices into the
  descriptor buffer binding points set by
  [vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBuffersEXT.html).

- `pOffsets` is a pointer to an array of
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) offsets to apply to the bound
  descriptor buffers.

## <a href="#_description" class="anchor"></a>Description

`vkCmdSetDescriptorBufferOffsetsEXT` binds `setCount` pairs of
descriptor buffers, specified by indices into the binding points bound
using
[vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBuffersEXT.html), and
buffer offsets to set numbers
\[`firstSet`..`firstSet`+`descriptorSetCount`-1\] for subsequent <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#pipelines-bindpoint-commands"
target="_blank" rel="noopener">bound pipeline commands</a> set by
`pipelineBindPoint`. Set \[`firstSet` + i\] is bound to the descriptor
buffer at binding `pBufferIndices`\[i\] at an offset of `pOffsets`\[i\].
Any bindings that were previously applied via these sets, or calls to
[vkCmdBindDescriptorSets](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorSets.html), are no longer
valid. Other sets will also be invalidated upon calling this command if
`layout` differs from the pipeline layout used to bind those other sets,
as described in <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#descriptorsets-compatibility"
target="_blank" rel="noopener">Pipeline Layout Compatibility</a>.

After binding descriptors, applications **can** modify descriptor memory
either by performing writes on the host or with device commands. When
descriptor memory is updated with device commands, visibility for the
shader stage accessing a descriptor is ensured with the
`VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT_EXT` access flag.
Implementations **must** not access resources referenced by these
descriptors unless they are dynamically accessed by shaders. Descriptors
bound with this call **can** be undefined if they are not dynamically
accessed by shaders.

Implementations **may** read descriptor data for any statically accessed
descriptor if the `binding` in `layout` is not declared with the
`VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT` flag. If the
`binding` in `layout` is declared with
`VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT`, implementations
**must** not read descriptor data that is not dynamically accessed.

Applications **must** ensure that any descriptor which the
implementation **may** read **must** be in-bounds of the underlying
descriptor buffer binding.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Applications can freely decide how large a variable descriptor buffer
binding is, so it may not be safe to read such descriptor payloads
statically. The intention of these rules is to allow implementations to
speculatively prefetch descriptor payloads where feasible.</p></td>
</tr>
</tbody>
</table>

Dynamically accessing a resource through descriptor data from an unbound
region of a <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#sparsememory-partially-resident-buffers"
target="_blank" rel="noopener">sparse partially-resident buffer</a> will
result in invalid descriptor data being read, and therefore undefined
behavior.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>For descriptors written by the host, visibility is implied through
the automatic visibility operation on queue submit, and there is no need
to consider <code>VK_ACCESS_2_DESCRIPTOR_BUFFER_READ_BIT</code>.
Explicit synchronization for descriptors is only required when
descriptors are updated on the device.</p></td>
</tr>
</tbody>
</table>

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>The requirements above imply that all descriptor bindings have been
defined with the equivalent of
<code>VK_DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT</code>,
<code>VK_DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT</code> and
<code>VK_DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT</code>, but enabling
those features is not required to get this behavior.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08061"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08061"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08061  
  The offsets in `pOffsets` **must** be aligned to
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`descriptorBufferOffsetAlignment`

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08063"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08063"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08063  
  The offsets in `pOffsets` **must** be small enough such that any
  descriptor binding referenced by `layout` without the
  `VK_DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT` flag computes a
  valid address inside the underlying [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html)

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08126"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08126"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08126  
  The offsets in `pOffsets` **must** be small enough such that any
  location accessed by a shader as a sampler descriptor **must** be
  within
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxSamplerDescriptorBufferRange`
  of the sampler descriptor buffer binding

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08127"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08127"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-08127  
  The offsets in `pOffsets` **must** be small enough such that any
  location accessed by a shader as a resource descriptor **must** be
  within
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxResourceDescriptorBufferRange`
  of the resource descriptor buffer binding

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-08064"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-08064"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-08064  
  Each element of `pBufferIndices` **must** be less than
  [VkPhysicalDeviceDescriptorBufferPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDescriptorBufferPropertiesEXT.html)::`maxDescriptorBufferBindings`

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-08065"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-08065"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-08065  
  Each element of `pBufferIndices` **must** reference a valid descriptor
  buffer binding set by a previous call to
  [vkCmdBindDescriptorBuffersEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdBindDescriptorBuffersEXT.html) in
  `commandBuffer`

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-firstSet-08066"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-firstSet-08066"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-firstSet-08066  
  The sum of `firstSet` and `setCount` **must** be less than or equal to
  [VkPipelineLayoutCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayoutCreateInfo.html)::`setLayoutCount`
  provided when `layout` was created

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-firstSet-09006"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-firstSet-09006"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-firstSet-09006  
  The [VkDescriptorSetLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDescriptorSetLayout.html) for each set
  from `firstSet` to `firstSet` + `setCount` when `layout` was created
  **must** have been created with the
  `VK_DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT` bit set

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-None-08060"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-None-08060"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-None-08060  
  The <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-descriptorBuffer"
  target="_blank" rel="noopener"><code>descriptorBuffer</code></a>
  feature **must** be enabled

- <a
  href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pipelineBindPoint-08067"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pipelineBindPoint-08067"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pipelineBindPoint-08067  
  `pipelineBindPoint` **must** be supported by the `commandBuffer`’s
  parent `VkCommandPool`’s queue family

Valid Usage (Implicit)

- <a
  href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-parameter"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-parameter"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a
  href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pipelineBindPoint-parameter"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pipelineBindPoint-parameter"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pipelineBindPoint-parameter  
  `pipelineBindPoint` **must** be a valid
  [VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html) value

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-layout-parameter"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-layout-parameter"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-layout-parameter  
  `layout` **must** be a valid [VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)
  handle

- <a
  href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-parameter"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-parameter"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pBufferIndices-parameter  
  `pBufferIndices` **must** be a valid pointer to an array of `setCount`
  `uint32_t` values

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-parameter"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-parameter"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-pOffsets-parameter  
  `pOffsets` **must** be a valid pointer to an array of `setCount`
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a
  href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-recording"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-recording"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-cmdpool"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-cmdpool"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics, or compute operations

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-videocoding"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-videocoding"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-setCount-arraylength"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-setCount-arraylength"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-setCount-arraylength  
  `setCount` **must** be greater than `0`

- <a href="#VUID-vkCmdSetDescriptorBufferOffsetsEXT-commonparent"
  id="VUID-vkCmdSetDescriptorBufferOffsetsEXT-commonparent"></a>
  VUID-vkCmdSetDescriptorBufferOffsetsEXT-commonparent  
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

[VK_EXT_descriptor_buffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_descriptor_buffer.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html),
[VkPipelineBindPoint](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineBindPoint.html),
[VkPipelineLayout](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPipelineLayout.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdSetDescriptorBufferOffsetsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
