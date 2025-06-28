# VkDeviceAddressBindingTypeEXT(3) Manual Page

## Name

VkDeviceAddressBindingTypeEXT - Enum describing a change in device address bindings



## [](#_c_specification)C Specification

The `VkDeviceAddressBindingTypeEXT` enum is defined as:

```c++
// Provided by VK_EXT_device_address_binding_report
typedef enum VkDeviceAddressBindingTypeEXT {
    VK_DEVICE_ADDRESS_BINDING_TYPE_BIND_EXT = 0,
    VK_DEVICE_ADDRESS_BINDING_TYPE_UNBIND_EXT = 1,
} VkDeviceAddressBindingTypeEXT;
```

## [](#_description)Description

- `VK_DEVICE_ADDRESS_BINDING_TYPE_BIND_EXT` specifies that a new GPU-accessible virtual address range has been bound.
- `VK_DEVICE_ADDRESS_BINDING_TYPE_UNBIND_EXT` specifies that a GPU-accessible virtual address range has been unbound.

## [](#_see_also)See Also

[VK\_EXT\_device\_address\_binding\_report](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_device_address_binding_report.html), [VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDeviceAddressBindingCallbackDataEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDeviceAddressBindingTypeEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0