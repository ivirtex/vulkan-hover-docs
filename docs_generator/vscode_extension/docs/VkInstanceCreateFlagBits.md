# VkInstanceCreateFlagBits(3) Manual Page

## Name

VkInstanceCreateFlagBits - Bitmask specifying behavior of the instance



## [](#_c_specification)C Specification

```c++
// Provided by VK_VERSION_1_0
typedef enum VkInstanceCreateFlagBits {
  // Provided by VK_KHR_portability_enumeration
    VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR = 0x00000001,
} VkInstanceCreateFlagBits;
```

## [](#_description)Description

- `VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR` specifies that the instance will enumerate available Vulkan Portability-compliant physical devices and groups in addition to the Vulkan physical devices and groups that are enumerated by default.

## [](#_see_also)See Also

[VK\_VERSION\_1\_0](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_VERSION_1_0.html), [VkInstanceCreateFlags](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateFlags.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkInstanceCreateFlagBits)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0