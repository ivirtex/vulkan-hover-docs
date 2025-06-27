# vkCmdBindIndexBuffer(3) Manual Page

## Name

vkCmdBindIndexBuffer - Bind an index buffer to a command buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

To bind an index buffer to a command buffer, call:

``` c
// Provided by VK_VERSION_1_0
void vkCmdBindIndexBuffer(
    VkCommandBuffer                             commandBuffer,
    VkBuffer                                    buffer,
    VkDeviceSize                                offset,
    VkIndexType                                 indexType);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `commandBuffer` is the command buffer into which the command is
  recorded.

- `buffer` is the buffer being bound.

- `offset` is the starting offset in bytes within `buffer` used in index
  buffer address calculations.

- `indexType` is a [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) value specifying the
  size of the indices.

## <a href="#_description" class="anchor"></a>Description

If the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-maintenance6"
target="_blank" rel="noopener"><code>maintenance6</code></a> feature is
enabled, `buffer` **can** be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html). If
`buffer` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html) and the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#features-nullDescriptor"
target="_blank" rel="noopener"><code>nullDescriptor</code></a> feature
is enabled, every index fetched results in a value of zero.

Valid Usage

- <a href="#VUID-vkCmdBindIndexBuffer-offset-08782"
  id="VUID-vkCmdBindIndexBuffer-offset-08782"></a>
  VUID-vkCmdBindIndexBuffer-offset-08782  
  `offset` **must** be less than the size of `buffer`

- <a href="#VUID-vkCmdBindIndexBuffer-offset-08783"
  id="VUID-vkCmdBindIndexBuffer-offset-08783"></a>
  VUID-vkCmdBindIndexBuffer-offset-08783  
  The sum of `offset` and the base address of the range of
  `VkDeviceMemory` object that is backing `buffer`, **must** be a
  multiple of the size of the type indicated by `indexType`

- <a href="#VUID-vkCmdBindIndexBuffer-buffer-08784"
  id="VUID-vkCmdBindIndexBuffer-buffer-08784"></a>
  VUID-vkCmdBindIndexBuffer-buffer-08784  
  `buffer` **must** have been created with the
  `VK_BUFFER_USAGE_INDEX_BUFFER_BIT` flag

- <a href="#VUID-vkCmdBindIndexBuffer-buffer-08785"
  id="VUID-vkCmdBindIndexBuffer-buffer-08785"></a>
  VUID-vkCmdBindIndexBuffer-buffer-08785  
  If `buffer` is non-sparse then it **must** be bound completely and
  contiguously to a single `VkDeviceMemory` object

- <a href="#VUID-vkCmdBindIndexBuffer-indexType-08786"
  id="VUID-vkCmdBindIndexBuffer-indexType-08786"></a>
  VUID-vkCmdBindIndexBuffer-indexType-08786  
  `indexType` **must** not be `VK_INDEX_TYPE_NONE_KHR`

- <a href="#VUID-vkCmdBindIndexBuffer-indexType-08787"
  id="VUID-vkCmdBindIndexBuffer-indexType-08787"></a>
  VUID-vkCmdBindIndexBuffer-indexType-08787  
  If `indexType` is `VK_INDEX_TYPE_UINT8_KHR`, the
  [`indexTypeUint8`](#features-indexTypeUint8) feature **must** be
  enabled

- <a href="#VUID-vkCmdBindIndexBuffer-None-09493"
  id="VUID-vkCmdBindIndexBuffer-None-09493"></a>
  VUID-vkCmdBindIndexBuffer-None-09493  
  If [`maintenance6`](#features-maintenance6) is not enabled, `buffer`
  **must** not be [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html)

- <a href="#VUID-vkCmdBindIndexBuffer-buffer-09494"
  id="VUID-vkCmdBindIndexBuffer-buffer-09494"></a>
  VUID-vkCmdBindIndexBuffer-buffer-09494  
  If `buffer` is [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), offset **must**
  be zero

Valid Usage (Implicit)

- <a href="#VUID-vkCmdBindIndexBuffer-commandBuffer-parameter"
  id="VUID-vkCmdBindIndexBuffer-commandBuffer-parameter"></a>
  VUID-vkCmdBindIndexBuffer-commandBuffer-parameter  
  `commandBuffer` **must** be a valid
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handle

- <a href="#VUID-vkCmdBindIndexBuffer-buffer-parameter"
  id="VUID-vkCmdBindIndexBuffer-buffer-parameter"></a>
  VUID-vkCmdBindIndexBuffer-buffer-parameter  
  If `buffer` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `buffer`
  **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBuffer.html) handle

- <a href="#VUID-vkCmdBindIndexBuffer-indexType-parameter"
  id="VUID-vkCmdBindIndexBuffer-indexType-parameter"></a>
  VUID-vkCmdBindIndexBuffer-indexType-parameter  
  `indexType` **must** be a valid [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html) value

- <a href="#VUID-vkCmdBindIndexBuffer-commandBuffer-recording"
  id="VUID-vkCmdBindIndexBuffer-commandBuffer-recording"></a>
  VUID-vkCmdBindIndexBuffer-commandBuffer-recording  
  `commandBuffer` **must** be in the [recording
  state](#commandbuffers-lifecycle)

- <a href="#VUID-vkCmdBindIndexBuffer-commandBuffer-cmdpool"
  id="VUID-vkCmdBindIndexBuffer-commandBuffer-cmdpool"></a>
  VUID-vkCmdBindIndexBuffer-commandBuffer-cmdpool  
  The `VkCommandPool` that `commandBuffer` was allocated from **must**
  support graphics operations

- <a href="#VUID-vkCmdBindIndexBuffer-videocoding"
  id="VUID-vkCmdBindIndexBuffer-videocoding"></a>
  VUID-vkCmdBindIndexBuffer-videocoding  
  This command **must** only be called outside of a video coding scope

- <a href="#VUID-vkCmdBindIndexBuffer-commonparent"
  id="VUID-vkCmdBindIndexBuffer-commonparent"></a>
  VUID-vkCmdBindIndexBuffer-commonparent  
  Both of `buffer`, and `commandBuffer` that are valid handles of
  non-ignored parameters **must** have been created, allocated, or
  retrieved from the same [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

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
[VkDeviceSize](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceSize.html), [VkIndexType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkIndexType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkCmdBindIndexBuffer"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
