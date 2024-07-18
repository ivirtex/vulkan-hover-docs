# vkAllocateCommandBuffers(3) Manual Page

## Name

vkAllocateCommandBuffers - Allocate command buffers from an existing
command pool



## <a href="#_c_specification" class="anchor"></a>C Specification

To allocate command buffers, call:

``` c
// Provided by VK_VERSION_1_0
VkResult vkAllocateCommandBuffers(
    VkDevice                                    device,
    const VkCommandBufferAllocateInfo*          pAllocateInfo,
    VkCommandBuffer*                            pCommandBuffers);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the command pool.

- `pAllocateInfo` is a pointer to a
  [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferAllocateInfo.html)
  structure describing parameters of the allocation.

- `pCommandBuffers` is a pointer to an array of
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handles in which the resulting
  command buffer objects are returned. The array **must** be at least
  the length specified by the `commandBufferCount` member of
  `pAllocateInfo`. Each allocated command buffer begins in the initial
  state.

## <a href="#_description" class="anchor"></a>Description

`vkAllocateCommandBuffers` **can** be used to allocate multiple command
buffers. If the allocation of any of those command buffers fails, the
implementation **must** free all successfully allocated command buffer
objects from this command, set all entries of the `pCommandBuffers`
array to `NULL` and return the error.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Filling <code>pCommandBuffers</code> with <code>NULL</code> values on
failure is an exception to the default error behavior that output
parameters will have undefined contents.</p></td>
</tr>
</tbody>
</table>

When command buffers are first allocated, they are in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#commandbuffers-lifecycle"
target="_blank" rel="noopener">initial state</a>.

Valid Usage (Implicit)

- <a href="#VUID-vkAllocateCommandBuffers-device-parameter"
  id="VUID-vkAllocateCommandBuffers-device-parameter"></a>
  VUID-vkAllocateCommandBuffers-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkAllocateCommandBuffers-pAllocateInfo-parameter"
  id="VUID-vkAllocateCommandBuffers-pAllocateInfo-parameter"></a>
  VUID-vkAllocateCommandBuffers-pAllocateInfo-parameter  
  `pAllocateInfo` **must** be a valid pointer to a valid
  [VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferAllocateInfo.html)
  structure

- <a href="#VUID-vkAllocateCommandBuffers-pCommandBuffers-parameter"
  id="VUID-vkAllocateCommandBuffers-pCommandBuffers-parameter"></a>
  VUID-vkAllocateCommandBuffers-pCommandBuffers-parameter  
  `pCommandBuffers` **must** be a valid pointer to an array of
  `pAllocateInfo->commandBufferCount`
  [VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html) handles

- <a
  href="#VUID-vkAllocateCommandBuffers-pAllocateInfo::commandBufferCount-arraylength"
  id="VUID-vkAllocateCommandBuffers-pAllocateInfo::commandBufferCount-arraylength"></a>
  VUID-vkAllocateCommandBuffers-pAllocateInfo::commandBufferCount-arraylength  
  `pAllocateInfo->commandBufferCount` **must** be greater than `0`

Host Synchronization

- Host access to `pAllocateInfo->commandPool` **must** be externally
  synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkCommandBuffer](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBuffer.html),
[VkCommandBufferAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCommandBufferAllocateInfo.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkAllocateCommandBuffers"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
