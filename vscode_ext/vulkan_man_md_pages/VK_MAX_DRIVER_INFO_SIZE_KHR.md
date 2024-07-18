# VK_MAX_DRIVER_INFO_SIZE(3) Manual Page

## Name

VK_MAX_DRIVER_INFO_SIZE - Length of a physical device driver information
string



## <a href="#_c_specification" class="anchor"></a>C Specification

`VK_MAX_DRIVER_INFO_SIZE` is the length in `char` values of an array
containing a driver information string, as returned in
[VkPhysicalDeviceDriverProperties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceDriverProperties.html)::`driverInfo`.

``` c
#define VK_MAX_DRIVER_INFO_SIZE           256U
```

or the equivalent

``` c
#define VK_MAX_DRIVER_INFO_SIZE_KHR       VK_MAX_DRIVER_INFO_SIZE
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_driver_properties](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_driver_properties.html),
[VK_VERSION_1_2](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_2.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_MAX_DRIVER_INFO_SIZE"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
