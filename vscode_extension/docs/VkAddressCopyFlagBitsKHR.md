# VkAddressCopyFlagBitsKHR(3) Manual Page

## Name

VkAddressCopyFlagBitsKHR - Bitmask specifying address copy parameters



## [](#_c_specification)C Specification

Bits which **can** be set in a [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html) mask are:

```c++
// Provided by VK_KHR_copy_memory_indirect
typedef enum VkAddressCopyFlagBitsKHR {
    VK_ADDRESS_COPY_DEVICE_LOCAL_BIT_KHR = 0x00000001,
    VK_ADDRESS_COPY_SPARSE_BIT_KHR = 0x00000002,
    VK_ADDRESS_COPY_PROTECTED_BIT_KHR = 0x00000004,
} VkAddressCopyFlagBitsKHR;
```

## [](#_description)Description

- `VK_ADDRESS_COPY_DEVICE_LOCAL_BIT_KHR` specifies that the address range is expected to be resident in device local memory. Specifying this value is optional, but **may** lead to improved performance if set accurately.
- `VK_ADDRESS_COPY_PROTECTED_BIT_KHR` specifies that the address range is allocated from protected memory.
- `VK_ADDRESS_COPY_SPARSE_BIT_KHR` specifies that the address range may not be fully bound to physical memory when accessed.

## [](#_see_also)See Also

[VK\_KHR\_copy\_memory\_indirect](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_copy_memory_indirect.html), [VkAddressCopyFlagsKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkAddressCopyFlagsKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkAddressCopyFlagBitsKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0