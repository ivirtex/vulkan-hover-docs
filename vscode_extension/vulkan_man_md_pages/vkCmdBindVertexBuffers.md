# vkCmdBindVertexBuffers(3) Manual Page

## Name

vkCmdBindVertexBuffers - Bind vertex buffers to a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To bind vertex buffers to a command buffer for use in subsequent drawing
commands, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdBindVertexBuffers(
    VkCommandBuffer                             commandBuffer,
    uint32_t                                    firstBinding,
    uint32_t                                    bindingCount,
    const VkBuffer*                             pBuffers,
    const VkDeviceSize*                         pOffsets);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `firstBinding` is the index of the first vertex input binding whose
  state is updated by the command.

- `bindingCount` is the number of vertex input bindings whose state is
  updated by the command.

- `pBuffers` is a pointer to an array of buffer handles.

- `pOffsets` is a pointer to an array of buffer offsets.

## <a href="#_description" class="anchor"></a>Description

The values taken from elements i of `pBuffers` and `pOffsets` replace
the current state for the vertex input binding `firstBinding` + i, for i
in \[0, `bindingCount`). The vertex input binding is updated to start at
the offset indicated by `pOffsets`\[i\] from the start of the buffer
`pBuffers`\[i\]. All vertex input attributes that use each of these
bindings will use these updated addresses in their address calculations
for subsequent drawing commands. If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
is enabled, elements of `pBuffers` **can** be
[VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), and **can** be used by the vertex
shader. If a vertex input attribute is bound to a vertex input binding
that is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), the values taken from
memory are considered to be zero, and missing G, B, or A components are
[filled with (0,0,1)](#fxvertex-input-extraction).

Valid Usage

- <a href="#VUID-vkCmdBindVertexBuffers-firstBinding-00624"
  id="VUID-vkCmdBindVertexBuffers-firstBinding-00624"></a>
  VUID-vkCmdBindVertexBuffers-firstBinding-00624  
  `firstBinding` **must** be less than
  `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-vkCmdBindVertexBuffers-firstBinding-00625"
  id="VUID-vkCmdBindVertexBuffers-firstBinding-00625"></a>
  VUID-vkCmdBindVertexBuffers-firstBinding-00625  
  The sum of `firstBinding` and `bindingCount` **must** be less than or
  equal to `VkPhysicalDeviceLimits`::`maxVertexInputBindings`

- <a href="#VUID-vkCmdBindVertexBuffers-pOffsets-00626"
  id="VUID-vkCmdBindVertexBuffers-pOffsets-00626"></a>
  VUID-vkCmdBindVertexBuffers-pOffsets-00626  
  All elements of `pOffsets` **must** be less than the size of the
  corresponding element in `pBuffers`

- <a href="#VUID-vkCmdBindVertexBuffers-pBuffers-00627"
  id="VUID-vkCmdBindVertexBuffers-pBuffers-00627"></a>
  VUID-vkCmdBindVertexBuffers-pBuffers-00627  
  All elements of `pBuffers` **must** have been created with the
  `VK_BUFFER_USAGE_VERTEX_BUFFER_BIT` flag

- <a href="#VUID-vkCmdBindVertexBuffers-pBuffers-00628"
  id="VUID-vkCmdBindVertexBuffers-pBuffers-00628"></a>
  VUID-vkCmdBindVertexBuffers-pBuffers-00628  
  Each element of `pBuffers` that is non-sparse **must** be bound
  completely and contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdBindVertexBuffers-pBuffers-04001"
  id="VUID-vkCmdBindVertexBuffers-pBuffers-04001"></a>
  VUID-vkCmdBindVertexBuffers-pBuffers-04001  
  If the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
  target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
  is not enabled, all elements of `pBuffers` **must** not be
  [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindVertexBuffers-pBuffers-04002"
  id="VUID-vkCmdBindVertexBuffers-pBuffers-04002"></a>
  VUID-vkCmdBindVertexBuffers-pBuffers-04002  
  If an element of `pBuffers` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html),
  then the corresponding element of `pOffsets` **must** be zero

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindVertexBuffers-commandBuffer-parameter"
  id="VUID-vkCmdBindVertexBuffers-commandBuffer-parameter"></a>
  VUID-vkCmdBindVertexBuffers-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindVertexBuffers-pBuffers-parameter"
  id="VUID-vkCmdBindVertexBuffers-pBuffers-parameter"></a>
  VUID-vkCmdBindVertexBuffers-pBuffers-parameter  
  `pBuffers` **must** be a valid pointer to an array of `bindingCount`
  valid or [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)
  [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handles

- <a href="#VUID-vkCmdBindVertexBuffers-pOffsets-parameter"
  id="VUID-vkCmdBindVertexBuffers-pOffsets-parameter"></a>
  VUID-vkCmdBindVertexBuffers-pOffsets-parameter  
  `pOffsets` **must** be a valid pointer to an array of `bindingCount`
  [VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html) values

- <a href="#VUID-vkCmdBindVertexBuffers-commandBuffer-recording"
  id="VUID-vkCmdBindVertexBuffers-commandBuffer-recording"></a>
  VUID-vkCmdBindVertexBuffers-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindVertexBuffers-commandBuffer-cmdpool"
  id="VUID-vkCmdBindVertexBuffers-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindVertexBuffers-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBindVertexBuffers-videocoding"
  id="VUID-vkCmdBindVertexBuffers-videocoding"></a>
  VUID-vkCmdBindVertexBuffers-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindVertexBuffers-bindingCount-arraylength"
  id="VUID-vkCmdBindVertexBuffers-bindingCount-arraylength"></a>
  VUID-vkCmdBindVertexBuffers-bindingCount-arraylength  
  `bindingCount` **must** be greater than `0`

- <a href="#VUID-vkCmdBindVertexBuffers-commonparent"
  id="VUID-vkCmdBindVertexBuffers-commonparent"></a>
  VUID-vkCmdBindVertexBuffers-commonparent  
  Both of `commandBuffer`, and the elements of `pBuffers` that are valid
  handles of non-ignored parameters **must** have been created,
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
<td class="tableblock halign-left valign-top"><p>Graphics</p></td>
<td class="tableblock halign-left valign-top"><p>State</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindVertexBuffers"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
