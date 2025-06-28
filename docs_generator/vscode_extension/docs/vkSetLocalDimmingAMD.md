# vkSetLocalDimmingAMD(3) Manual Page

## Name

vkSetLocalDimmingAMD - Set Local Dimming



## [](#_c_specification)C Specification

The local dimming HDR setting may also be changed over the life of a swapchain by calling:

```c++
// Provided by VK_AMD_display_native_hdr
void vkSetLocalDimmingAMD(
    VkDevice                                    device,
    VkSwapchainKHR                              swapChain,
    VkBool32                                    localDimmingEnable);
```

## [](#_parameters)Parameters

- `device` is the device associated with `swapChain`.
- `swapChain` handle to enable local dimming.
- `localDimmingEnable` specifies whether local dimming is enabled for the swapchain.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-vkSetLocalDimmingAMD-device-parameter)VUID-vkSetLocalDimmingAMD-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html) handle
- [](#VUID-vkSetLocalDimmingAMD-swapChain-parameter)VUID-vkSetLocalDimmingAMD-swapChain-parameter  
  `swapChain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html) handle
- [](#VUID-vkSetLocalDimmingAMD-swapChain-parent)VUID-vkSetLocalDimmingAMD-swapChain-parent  
  `swapChain` **must** have been created, allocated, or retrieved from `device`

Valid Usage

- [](#VUID-vkSetLocalDimmingAMD-localDimmingSupport-04618)VUID-vkSetLocalDimmingAMD-localDimmingSupport-04618  
  [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html)::`localDimmingSupport` **must** be supported

## [](#_see_also)See Also

[VK\_AMD\_display\_native\_hdr](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_AMD_display_native_hdr.html), [VkBool32](https://registry.khronos.org/vulkan/specs/latest/man/html/VkBool32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/latest/man/html/VkDevice.html), [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/latest/man/html/VkSwapchainKHR.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#vkSetLocalDimmingAMD)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0