# VkLayeredDriverUnderlyingApiMSFT(3) Manual Page

## Name

VkLayeredDriverUnderlyingApiMSFT - Layered driver underlying APIs



## <a href="#_c_specification" class="anchor"></a>C Specification

Underlying APIs which **may** be returned in
[VkPhysicalDeviceLayeredDriverPropertiesMSFT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredDriverPropertiesMSFT.html)::`underlyingAPI`
are:

``` c
// Provided by VK_MSFT_layered_driver
typedef enum VkLayeredDriverUnderlyingApiMSFT {
    VK_LAYERED_DRIVER_UNDERLYING_API_NONE_MSFT = 0,
    VK_LAYERED_DRIVER_UNDERLYING_API_D3D12_MSFT = 1,
} VkLayeredDriverUnderlyingApiMSFT;
```

## <a href="#_see_also" class="anchor"></a>See Also

[VK_MSFT_layered_driver](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_MSFT_layered_driver.html),
[VkPhysicalDeviceLayeredDriverPropertiesMSFT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredDriverPropertiesMSFT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkLayeredDriverUnderlyingApiMSFT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
