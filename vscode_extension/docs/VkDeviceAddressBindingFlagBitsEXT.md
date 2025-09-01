# VkDeviceAddressBindingFlagBitsEXT(3) Manual Page

## Name

VkDeviceAddressBindingFlagBitsEXT - Bitmask specifying the additional information about a binding event



## [](#_c_specification)C Specification

Bits which **can** be set in [VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingCallbackDataEXT.html)::`flags` specifying additional information about a binding event are:

```c++
// Provided by VK_EXT_device_address_binding_report
typedef enum VkDeviceAddressBindingFlagBitsEXT {
    VK_DEVICE_ADDRESS_BINDING_INTERNAL_OBJECT_BIT_EXT = 0x00000001,
} VkDeviceAddressBindingFlagBitsEXT;
```

## [](#_description)Description

- `VK_DEVICE_ADDRESS_BINDING_INTERNAL_OBJECT_BIT_EXT` specifies that [VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingCallbackDataEXT.html) describes a Vulkan object that has not been made visible to the application via a Vulkan command.

## [](#_see_also)See Also

[VK\_EXT\_device\_address\_binding\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_address_binding_report.html), [VkDeviceAddressBindingFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceAddressBindingFlagBitsEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0