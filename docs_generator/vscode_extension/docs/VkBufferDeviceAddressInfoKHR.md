# VkBufferDeviceAddressInfo(3) Manual Page

## Name

VkBufferDeviceAddressInfo - Structure specifying the buffer to query an address for



## [](#_c_specification)C Specification

The `VkBufferDeviceAddressInfo` structure is defined as:

```c++
// Provided by VK_VERSION_1_2
typedef struct VkBufferDeviceAddressInfo {
    VkStructureType    sType;
    const void*        pNext;
    VkBuffer           buffer;
} VkBufferDeviceAddressInfo;
```

or the equivalent

```c++
// Provided by VK_KHR_buffer_device_address
typedef VkBufferDeviceAddressInfo VkBufferDeviceAddressInfoKHR;
```

or the equivalent

```c++
// Provided by VK_EXT_buffer_device_address
typedef VkBufferDeviceAddressInfo VkBufferDeviceAddressInfoEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `buffer` specifies the buffer whose address is being queried.

## [](#_description)Description

Valid Usage

- [](#VUID-VkBufferDeviceAddressInfo-buffer-02601)VUID-VkBufferDeviceAddressInfo-buffer-02601  
  `buffer` **must** have been created with `VK_BUFFER_USAGE_SHADER_DEVICE_ADDRESS_BIT`

Valid Usage (Implicit)

- [](#VUID-VkBufferDeviceAddressInfo-sType-sType)VUID-VkBufferDeviceAddressInfo-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_BUFFER_DEVICE_ADDRESS_INFO`
- [](#VUID-VkBufferDeviceAddressInfo-pNext-pNext)VUID-VkBufferDeviceAddressInfo-pNext-pNext  
  `pNext` **must** be `NULL`
- [](#VUID-VkBufferDeviceAddressInfo-buffer-parameter)VUID-VkBufferDeviceAddressInfo-buffer-parameter  
  `buffer` **must** be a valid [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html) handle

## [](#_see_also)See Also

[VK\_EXT\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_buffer_device_address.html), [VK\_KHR\_buffer\_device\_address](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_buffer_device_address.html), [VK\_VERSION\_1\_2](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_2.html), [VkBuffer](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBuffer.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [vkGetBufferDeviceAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddress.html), [vkGetBufferDeviceAddressEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddressEXT.html), [vkGetBufferDeviceAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferDeviceAddressKHR.html), [vkGetBufferOpaqueCaptureAddress](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureAddress.html), [vkGetBufferOpaqueCaptureAddressKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/vkGetBufferOpaqueCaptureAddressKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkBufferDeviceAddressInfo)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0