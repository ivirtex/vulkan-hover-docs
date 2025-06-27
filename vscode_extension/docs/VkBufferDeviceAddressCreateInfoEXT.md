# VkBufferDeviceAddressCreateInfoEXT(3) Manual Page

## Name

VkBufferDeviceAddressCreateInfoEXT - Request a specific address for a
buffer



## <a href="#_c_specification" class="anchor"></a>C Specification

Alternatively, to request a specific device address for a buffer, add a
[VkBufferDeviceAddressCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferDeviceAddressCreateInfoEXT.html)
structure to the `pNext` chain of the
[VkBufferCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBufferCreateInfo.html) structure. The
`VkBufferDeviceAddressCreateInfoEXT` structure is defined as:

``` c
// Provided by VK_EXT_buffer_device_address
typedef struct VkBufferDeviceAddressCreateInfoEXT {
    VkStructureType    sType;
    const void*        pNext;
    VkDeviceAddress    deviceAddress;
} VkBufferDeviceAddressCreateInfoEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `deviceAddress` is the device address requested for the buffer.

## <a href="#_description" class="anchor"></a>Description

If `deviceAddress` is zero, no specific address is requested.

If `deviceAddress` is not zero, then it **must** be an address retrieved
from an identically created buffer on the same implementation. The
buffer **must** also be bound to an identically created `VkDeviceMemory`
object.

If this structure is not present, it is as if `deviceAddress` is zero.

Apps **should** avoid creating buffers with app-provided addresses and
implementation-provided addresses in the same process, to reduce the
likelihood of `VK_ERROR_INVALID_DEVICE_ADDRESS_EXT` errors.

Valid Usage (Implicit)

- <a href="#VUID-VkBufferDeviceAddressCreateInfoEXT-sType-sType"
  id="VUID-VkBufferDeviceAddressCreateInfoEXT-sType-sType"></a>
  VUID-VkBufferDeviceAddressCreateInfoEXT-sType-sType  
  `sType` **must** be
  `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_CREATE_INFO_EXT`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_buffer_device_address](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_buffer_device_address.html),
[VkDeviceAddress](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddress.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkBufferDeviceAddressCreateInfoEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
