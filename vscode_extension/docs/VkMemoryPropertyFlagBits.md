# VkMemoryPropertyFlagBits(3) Manual Page

## Name

VkMemoryPropertyFlagBits - Bitmask specifying properties for a memory
type



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **may** be set in
[VkMemoryType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryType.html)::`propertyFlags`, indicating
properties of a memory type, are:

``` c
// Provided by VK_VERSION_1_0
typedef enum VkMemoryPropertyFlagBits {
    VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT = 0x00000001,
    VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT = 0x00000002,
    VK_MEMORY_PROPERTY_HOST_COHERENT_BIT = 0x00000004,
    VK_MEMORY_PROPERTY_HOST_CACHED_BIT = 0x00000008,
    VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT = 0x00000010,
  // Provided by VK_VERSION_1_1
    VK_MEMORY_PROPERTY_PROTECTED_BIT = 0x00000020,
  // Provided by VK_AMD_device_coherent_memory
    VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD = 0x00000040,
  // Provided by VK_AMD_device_coherent_memory
    VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD = 0x00000080,
  // Provided by VK_NV_external_memory_rdma
    VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV = 0x00000100,
} VkMemoryPropertyFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` bit specifies that memory
  allocated with this type is the most efficient for device access. This
  property will be set if and only if the memory type belongs to a heap
  with the `VK_MEMORY_HEAP_DEVICE_LOCAL_BIT` set.

- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` bit specifies that memory
  allocated with this type **can** be mapped for host access using
  [vkMapMemory](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkMapMemory.html).

- <span id="memory-coherent"></span>
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` bit specifies that the host
  cache management commands
  [vkFlushMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkFlushMappedMemoryRanges.html) and
  [vkInvalidateMappedMemoryRanges](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkInvalidateMappedMemoryRanges.html)
  are not needed to flush host writes to the device or make device
  writes visible to the host, respectively.

- `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` bit specifies that memory
  allocated with this type is cached on the host. Host memory accesses
  to uncached memory are slower than to cached memory, however uncached
  memory is always host coherent.

- `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT` bit specifies that the
  memory type only allows device access to the memory. Memory types
  **must** not have both `VK_MEMORY_PROPERTY_LAZILY_ALLOCATED_BIT` and
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` set. Additionally, the object’s
  backing memory **may** be provided by the implementation lazily as
  specified in <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#memory-device-lazy_allocation"
  target="_blank" rel="noopener">Lazily Allocated Memory</a>.

- `VK_MEMORY_PROPERTY_PROTECTED_BIT` bit specifies that the memory type
  only allows device access to the memory, and allows protected queue
  operations to access the memory. Memory types **must** not have
  `VK_MEMORY_PROPERTY_PROTECTED_BIT` set and any of
  `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` set, or
  `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` set, or
  `VK_MEMORY_PROPERTY_HOST_CACHED_BIT` set.

- `VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD` bit specifies that device
  accesses to allocations of this memory type are automatically made
  available and visible.

- `VK_MEMORY_PROPERTY_DEVICE_UNCACHED_BIT_AMD` bit specifies that memory
  allocated with this type is not cached on the device. Uncached device
  memory is always device coherent.

- `VK_MEMORY_PROPERTY_RDMA_CAPABLE_BIT_NV` bit specifies that external
  devices can access this memory directly.

For any memory allocated with both the
`VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` and the
`VK_MEMORY_PROPERTY_DEVICE_COHERENT_BIT_AMD`, host or device accesses
also perform automatic memory domain transfer operations, such that
writes are always automatically available and visible to both host and
device memory domains.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr>
<td class="icon"><em></em></td>
<td class="content">Note
<p>Device coherence is a useful property for certain debugging use cases
(e.g. crash analysis, where performing separate coherence actions could
mean values are not reported correctly). However, device coherent
accesses may be slower than equivalent accesses without device
coherence, particularly if they are also device uncached. For device
uncached memory in particular, repeated accesses to the same or
neighboring memory locations over a short time period (e.g. within a
frame) may be slower than it would be for the equivalent cached memory
type. As such, it is generally inadvisable to use device coherent or
device uncached memory except when really needed.</p></td>
</tr>
</tbody>
</table>

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkMemoryPropertyFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryPropertyFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryPropertyFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
