# VkDeviceAddressBindingFlagBitsEXT(3) Manual Page

## Name

VkDeviceAddressBindingFlagBitsEXT - Bitmask specifying the additional
information about a binding event



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which **can** be set in
[VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddressBindingCallbackDataEXT.html)::`flags`
specifying additional information about a binding event are:

``` c
// Provided by VK_EXT_device_address_binding_report
typedef enum VkDeviceAddressBindingFlagBitsEXT {
    VK_DEVICE_ADDRESS_BINDING_INTERNAL_OBJECT_BIT_EXT = 0x00000001,
} VkDeviceAddressBindingFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_DEVICE_ADDRESS_BINDING_INTERNAL_OBJECT_BIT_EXT` specifies that
  [VkDeviceAddressBindingCallbackDataEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddressBindingCallbackDataEXT.html)
  describes a Vulkan object that has not been made visible to the
  application via a Vulkan command.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_device_address_binding_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_device_address_binding_report.html),
[VkDeviceAddressBindingFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceAddressBindingFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkDeviceAddressBindingFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
