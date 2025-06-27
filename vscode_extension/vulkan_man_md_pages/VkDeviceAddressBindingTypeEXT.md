# VkDeviceAddressBindingTypeEXT(3) Manual Page

## Name

VkDeviceAddressBindingTypeEXT - Enum describing a change in device
address bindings



## <a href="#_c_specification" class="anchor"></a>C Specification

The `VkDeviceAddressBindingTypeEXT` enum is defined as:

``` c
// Provided by VK_EXT_device_address_binding_report
typedef enum VkDeviceAddressBindingTypeEXT {
    VK_DEVICE_ADDRESS_BINDING_TYPE_BIND_EXT = 0,
    VK_DEVICE_ADDRESS_BINDING_TYPE_UNBIND_EXT = 1,
} VkDeviceAddressBindingTypeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_ADDRESS_BINDING_TYPE_BIND_EXT` specifies that a new
  GPU-accessible virtual address range has been bound.

- `VK_DEVICE_ADDRESS_BINDING_TYPE_UNBIND_EXT` specifies that a
  GPU-accessible virtual address range has been unbound.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_address_binding_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_address_binding_report.html),
[VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddressBindingCallbackDataEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceAddressBindingTypeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
