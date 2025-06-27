# VkInstanceCreateFlagBits(3) Manual Page

## Name

VkInstanceCreateFlagBits - Bitmask specifying behavior of the instance



## <a href="#_c_specification" class="anchor"></a>C Specification

``` c
// Provided by VK_VERSION_1_0
typedef enum VkInstanceCreateFlagBits {
  // Provided by VK_KHR_portability_enumeration
    VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR = 0x00000001,
} VkInstanceCreateFlagBits;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_INSTANCE_CREATE_ENUMERATE_PORTABILITY_BIT_KHR` specifies that the
  instance will enumerate available Vulkan Portability-compliant
  physical devices and groups in addition to the Vulkan physical devices
  and groups that are enumerated by default.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_VERSION_1_0](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_0.html),
[VkInstanceCreateFlags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateFlags.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkInstanceCreateFlagBits"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
