# vkGetDeviceGroupSurfacePresentModesKHR(3) Manual Page

## Name

vkGetDeviceGroupSurfacePresentModesKHR - Query present capabilities for
a surface



## <a href="#_c_specification" class="anchor"></a>C Specification

Some surfaces **may** not be capable of using all the device group
present modes.

To query the supported device group present modes for a particular
surface, call:

``` c
// Provided by VK_VERSION_1_1 with VK_KHR_swapchain, VK_KHR_device_group with VK_KHR_surface
VkResult vkGetDeviceGroupSurfacePresentModesKHR(
    VkDevice                                    device,
    VkSurfaceKHR                                surface,
    VkDeviceGroupPresentModeFlagsKHR*           pModes);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the logical device.

- `surface` is the surface.

- `pModes` is a pointer to a
  [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html)
  in which the supported device group present modes for the surface are
  returned.

## <a href="#_description" class="anchor"></a>Description

The modes returned by this command are not invariant, and **may** change
in response to the surface being moved, resized, or occluded. These
modes **must** be a subset of the modes returned by
[vkGetDeviceGroupPresentCapabilitiesKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetDeviceGroupPresentCapabilitiesKHR.html).

Valid Usage

- <a href="#VUID-vkGetDeviceGroupSurfacePresentModesKHR-surface-06212"
  id="VUID-vkGetDeviceGroupSurfacePresentModesKHR-surface-06212"></a>
  VUID-vkGetDeviceGroupSurfacePresentModesKHR-surface-06212  
  `surface` **must** be supported by all physical devices associated
  with `device`, as reported by
  [vkGetPhysicalDeviceSurfaceSupportKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkGetPhysicalDeviceSurfaceSupportKHR.html)
  or an equivalent platform-specific mechanism

Valid Usage (Implicit)

- <a href="#VUID-vkGetDeviceGroupSurfacePresentModesKHR-device-parameter"
  id="VUID-vkGetDeviceGroupSurfacePresentModesKHR-device-parameter"></a>
  VUID-vkGetDeviceGroupSurfacePresentModesKHR-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkGetDeviceGroupSurfacePresentModesKHR-surface-parameter"
  id="VUID-vkGetDeviceGroupSurfacePresentModesKHR-surface-parameter"></a>
  VUID-vkGetDeviceGroupSurfacePresentModesKHR-surface-parameter  
  `surface` **must** be a valid [VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html) handle

- <a href="#VUID-vkGetDeviceGroupSurfacePresentModesKHR-pModes-parameter"
  id="VUID-vkGetDeviceGroupSurfacePresentModesKHR-pModes-parameter"></a>
  VUID-vkGetDeviceGroupSurfacePresentModesKHR-pModes-parameter  
  `pModes` **must** be a valid pointer to a
  [VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html)
  value

- <a href="#VUID-vkGetDeviceGroupSurfacePresentModesKHR-commonparent"
  id="VUID-vkGetDeviceGroupSurfacePresentModesKHR-commonparent"></a>
  VUID-vkGetDeviceGroupSurfacePresentModesKHR-commonparent  
  Both of `device`, and `surface` **must** have been created, allocated,
  or retrieved from the same [VkInstance](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstance.html)

Host Synchronization

- Host access to `surface` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

- `VK_ERROR_SURFACE_LOST_KHR`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_KHR_device_group](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_device_group.html),
[VK_KHR_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_surface.html),
[VK_KHR_swapchain](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_swapchain.html),
[VK_VERSION_1_1](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_VERSION_1_1.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkDeviceGroupPresentModeFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDeviceGroupPresentModeFlagsKHR.html),
[VkSurfaceKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSurfaceKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetDeviceGroupSurfacePresentModesKHR"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
