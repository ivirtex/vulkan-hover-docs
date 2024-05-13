# vkFreeMemory(3) Manual Page

## Name

vkFreeMemory - Free device memory



## <a href="#_c_specification" class="anchor"></a>C Specification

To free a memory object, call:

``` c
// Provided by VK_VERSION_1_0
void vkFreeMemory(
    VkDevice                                    device,
    VkDeviceMemory                              memory,
    const VkAllocationCallbacks*                pAllocator);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device that owns the memory.

- `memory` is the [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) object to be
  freed.

- `pAllocator` controls host memory allocation as described in the <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-allocation"
  target="_blank" rel="noopener">Memory Allocation</a> chapter.

## <a href="#_description" class="anchor"></a>Description

Before freeing a memory object, an application **must** ensure the
memory object is no longer in use by the device — for example by command
buffers in the *pending state*. Memory **can** be freed whilst still
bound to resources, but those resources **must** not be used afterwards.
Freeing a memory object releases the reference it held, if any, to its
payload. If there are still any bound images or buffers, the memory
object’s payload **may** not be immediately released by the
implementation, but **must** be released by the time all bound images
and buffers have been destroyed. Once all references to a payload are
released, it is returned to the heap from which it was allocated.

How memory objects are bound to Images and Buffers is described in
detail in the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#resources-association"
target="_blank" rel="noopener">Resource Memory Association</a> section.

If a memory object is mapped at the time it is freed, it is implicitly
unmapped.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>As described <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device-unmap-does-not-flush"
target="_blank" rel="noopener">below</a>, host writes are not implicitly
flushed when the memory object is unmapped, but the implementation
<strong>must</strong> guarantee that writes that have not been flushed
do not affect any other memory.</p></td>
</tr>
</tbody>
</table>

Valid Usage

- <a href="#VUID-vkFreeMemory-memory-00677"
  id="VUID-vkFreeMemory-memory-00677"></a>
  VUID-vkFreeMemory-memory-00677  
  All submitted commands that refer to `memory` (via images or buffers)
  **must** have completed execution

Valid Usage (Implicit)

- <a href="#VUID-vkFreeMemory-device-parameter"
  id="VUID-vkFreeMemory-device-parameter"></a>
  VUID-vkFreeMemory-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkFreeMemory-memory-parameter"
  id="VUID-vkFreeMemory-memory-parameter"></a>
  VUID-vkFreeMemory-memory-parameter  
  If `memory` is not [VK_NULL_HANDLE](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_NULL_HANDLE.html), `memory`
  **must** be a valid [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html) handle

- <a href="#VUID-vkFreeMemory-pAllocator-parameter"
  id="VUID-vkFreeMemory-pAllocator-parameter"></a>
  VUID-vkFreeMemory-pAllocator-parameter  
  If `pAllocator` is not `NULL`, `pAllocator` **must** be a valid
  pointer to a valid [VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html)
  structure

- <a href="#VUID-vkFreeMemory-memory-parent"
  id="VUID-vkFreeMemory-memory-parent"></a>
  VUID-vkFreeMemory-memory-parent  
  If `memory` is a valid handle, it **must** have been created,
  allocated, or retrieved from `device`

Host Synchronization

- Host access to `memory` **must** be externally synchronized

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkAllocationCallbacks](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkAllocationCallbacks.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDeviceMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceMemory.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkFreeMemory"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
