# VkDeviceFaultAddressInfoEXT(3) Manual Page

## Name

VkDeviceFaultAddressInfoEXT - Structure specifying GPU virtual address information



## [](#_c_specification)C Specification

The `VkDeviceFaultAddressInfoEXT` structure is defined as:

```c++
// Provided by VK_EXT_device_fault
typedef struct VkDeviceFaultAddressInfoEXT {
    VkDeviceFaultAddressTypeEXT    addressType;
    VkDeviceAddress                reportedAddress;
    VkDeviceSize                   addressPrecision;
} VkDeviceFaultAddressInfoEXT;
```

## [](#_members)Members

- `addressType` is either the type of memory operation that triggered a page fault, or the type of association between an instruction pointer and a fault.
- `reportedAddress` is the GPU virtual address recorded by the device.
- `addressPrecision` is a power of two value that specifies how precisely the device can report the address.

## [](#_description)Description

The combination of `reportedAddress` and `addressPrecision` allow the possible range of addresses to be calculated, such that:

```c++
lower_address = (pInfo->reportedAddress & ~(pInfo->addressPrecision-1))
upper_address = (pInfo->reportedAddress |  (pInfo->addressPrecision-1))
```

Note

It is valid for the `reportedAddress` to contain a more precise address than indicated by `addressPrecision`. In this case, the value of `reportedAddress` should be treated as an additional hint as to the value of the address that triggered the page fault, or to the value of an instruction pointer.

Valid Usage (Implicit)

- [](#VUID-VkDeviceFaultAddressInfoEXT-addressType-parameter)VUID-VkDeviceFaultAddressInfoEXT-addressType-parameter  
  `addressType` **must** be a valid [VkDeviceFaultAddressTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressTypeEXT.html) value
- [](#VUID-VkDeviceFaultAddressInfoEXT-reportedAddress-parameter)VUID-VkDeviceFaultAddressInfoEXT-reportedAddress-parameter  
  `reportedAddress` **must** be a valid [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html) value

## [](#_see_also)See Also

[VK\_EXT\_device\_fault](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_fault.html), [VkDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddress.html), [VkDeviceFaultAddressTypeEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultAddressTypeEXT.html), [VkDeviceFaultInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceFaultInfoEXT.html), [VkDeviceSize](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceSize.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceFaultAddressInfoEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0