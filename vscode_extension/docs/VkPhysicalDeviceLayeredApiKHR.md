# VkPhysicalDeviceLayeredApiKHR(3) Manual Page

## Name

VkPhysicalDeviceLayeredApiKHR - API implemented by the layered implementation



## [](#_c_specification)C Specification

The list of possible API implementations of a layered implementation underneath the Vulkan physical device, as returned in [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)::`layeredAPI`, are:

```c++
// Provided by VK_KHR_maintenance7
typedef enum VkPhysicalDeviceLayeredApiKHR {
    VK_PHYSICAL_DEVICE_LAYERED_API_VULKAN_KHR = 0,
    VK_PHYSICAL_DEVICE_LAYERED_API_D3D12_KHR = 1,
    VK_PHYSICAL_DEVICE_LAYERED_API_METAL_KHR = 2,
    VK_PHYSICAL_DEVICE_LAYERED_API_OPENGL_KHR = 3,
    VK_PHYSICAL_DEVICE_LAYERED_API_OPENGLES_KHR = 4,
} VkPhysicalDeviceLayeredApiKHR;
```

## [](#_description)Description

- `VK_PHYSICAL_DEVICE_LAYERED_API_VULKAN_KHR` - the device implements the Vulkan API.
- `VK_PHYSICAL_DEVICE_LAYERED_API_D3D12_KHR` - the device implements the D3D12 API.
- `VK_PHYSICAL_DEVICE_LAYERED_API_METAL_KHR` - the device implements the Metal API.
- `VK_PHYSICAL_DEVICE_LAYERED_API_OPENGL_KHR` - the device implements the OpenGL API.
- `VK_PHYSICAL_DEVICE_LAYERED_API_OPENGLES_KHR` - the device implements the OpenGL ES API.

## [](#_see_also)See Also

[VK\_KHR\_maintenance7](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_KHR_maintenance7.html), [VkPhysicalDeviceLayeredApiPropertiesKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkPhysicalDeviceLayeredApiPropertiesKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkPhysicalDeviceLayeredApiKHR).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0