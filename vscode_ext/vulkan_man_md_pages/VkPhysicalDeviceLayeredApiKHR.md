# VkPhysicalDeviceLayeredApiKHR(3) Manual Page

## Name

VkPhysicalDeviceLayeredApiKHR - API implemented by the layered
implementation



## <a href="#_c_specification" class="anchor"></a>C Specification

The list of possible API implementations of a layered implementation
underneath the Vulkan physical device, as returned in
[VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)::`layeredAPI`,
are:

``` c
// Provided by VK_KHR_maintenance7
typedef enum VkPhysicalDeviceLayeredApiKHR {
    VK_PHYSICAL_DEVICE_LAYERED_API_VULKAN_KHR = 0,
    VK_PHYSICAL_DEVICE_LAYERED_API_D3D12_KHR = 1,
    VK_PHYSICAL_DEVICE_LAYERED_API_METAL_KHR = 2,
    VK_PHYSICAL_DEVICE_LAYERED_API_OPENGL_KHR = 3,
    VK_PHYSICAL_DEVICE_LAYERED_API_OPENGLES_KHR = 4,
} VkPhysicalDeviceLayeredApiKHR;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_PHYSICAL_DEVICE_LAYERED_API_VULKAN_KHR` - the device implements
  the Vulkan API.

- `VK_PHYSICAL_DEVICE_LAYERED_API_D3D12_KHR` - the device implements the
  D3D12 API.

- `VK_PHYSICAL_DEVICE_LAYERED_API_METAL_KHR` - the device implements the
  Metal API.

- `VK_PHYSICAL_DEVICE_LAYERED_API_OPENGL_KHR` - the device implements
  the OpenGL API.

- `VK_PHYSICAL_DEVICE_LAYERED_API_OPENGLES_KHR` - the device implements
  the OpenGL ES API.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_maintenance7](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_maintenance7.html),
[VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkPhysicalDeviceLayeredApiKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
