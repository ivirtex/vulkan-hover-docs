# VkMemoryOpaqueCaptureAddressAllocateInfo(3) Manual Page

## Name

VkMemoryOpaqueCaptureAddressAllocateInfo - Request a specific address
for a memory allocation



## <a href="#_c_specification" class="anchor"></a>C Specification

To request a specific device address for a memory allocation, add a
[VkMemoryOpaqueCaptureAddressAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryOpaqueCaptureAddressAllocateInfo.html)
structure to the `pNext` chain of the
[VkMemoryAllocateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkMemoryAllocateInfo.html) structure. The
`VkMemoryOpaqueCaptureAddressAllocateInfo` structure is defined as:

``` c
// Provided by VK_VERSION_1_2
typedef struct VkMemoryOpaqueCaptureAddressAllocateInfo {
    VkStructureType    sType;
    const void*        pNext;
    uint64_t           opaqueCaptureAddress;
} VkMemoryOpaqueCaptureAddressAllocateInfo;
```

or the equivalent

``` c
// Provided by VK_KHR_buffer_device_address
typedef VkMemoryOpaqueCaptureAddressAllocateInfo VkMemoryOpaqueCaptureAddressAllocateInfoKHR;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `opaqueCaptureAddress` is the opaque capture address requested for the
  memory allocation.

## <a href="#_description" class="anchor"></a>Description

If `opaqueCaptureAddress` is zero, no specific address is requested.

If `opaqueCaptureAddress` is not zero, it **should** be an address
retrieved from
[vkGetDeviceMemoryOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html)
on an identically created memory allocation on the same implementation.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>In most cases, it is expected that a non-zero
<code>opaqueAddress</code> is an address retrieved from <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceMemoryOpaqueCaptureAddress.html">vkGetDeviceMemoryOpaqueCaptureAddress</a>
on an identically created memory allocation. If this is not the case, it
is likely that <code>VK_ERROR_INVALID_OPAQUE_CAPTURE_ADDRESS</code>
errors will occur.</p>
<p>This is, however, not a strict requirement because trace
capture/replay tools may need to adjust memory allocation parameters for
imported memory.</p></td>
</tr>
</tbody>
</table>

If this structure is not present, it is as if `opaqueCaptureAddress` is
zero.

Valid Usage (Implicit)

- <a href="#VUID-VkMemoryOpaqueCaptureAddressAllocateInfo-sType-sType"
  id="VUID-VkMemoryOpaqueCaptureAddressAllocateInfo-sType-sType"></a>
  VUID-VkMemoryOpaqueCaptureAddressAllocateInfo-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_MEMORY_OPAQUE_CAPTURE_ADDRESS_ALLOCATE_INFO`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_buffer_device_address.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkMemoryOpaqueCaptureAddressAllocateInfo"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
