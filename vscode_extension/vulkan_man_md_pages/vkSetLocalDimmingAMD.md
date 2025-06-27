# vkSetLocalDimmingAMD(3) Manual Page

## Name

vkSetLocalDimmingAMD - Set Local Dimming



## <a href="#_c_specification" class="anchor"></a>C Specification

The local dimming HDR setting may also be changed over the life of a
swapchain by calling:

``` c
// Provided by VK_AMD_display_native_hdr
void vkSetLocalDimmingAMD(
    VkDevice                                    device,
    VkSwapchainKHR                              swapChain,
    VkBool32                                    localDimmingEnable);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device associated with `swapChain`.

- `swapChain` handle to enable local dimming.

- `localDimmingEnable` specifies whether local dimming is enabled for
  the swapchain.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkSetLocalDimmingAMD-device-parameter"
  id="VUID-vkSetLocalDimmingAMD-device-parameter"></a>
  VUID-vkSetLocalDimmingAMD-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkSetLocalDimmingAMD-swapChain-parameter"
  id="VUID-vkSetLocalDimmingAMD-swapChain-parameter"></a>
  VUID-vkSetLocalDimmingAMD-swapChain-parameter  
  `swapChain` **must** be a valid [VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)
  handle

- <a href="#VUID-vkSetLocalDimmingAMD-swapChain-parent"
  id="VUID-vkSetLocalDimmingAMD-swapChain-parent"></a>
  VUID-vkSetLocalDimmingAMD-swapChain-parent  
  `swapChain` **must** have been created, allocated, or retrieved from
  `device`

Valid Usage

- <a href="#VUID-vkSetLocalDimmingAMD-localDimmingSupport-04618"
  id="VUID-vkSetLocalDimmingAMD-localDimmingSupport-04618"></a>
  VUID-vkSetLocalDimmingAMD-localDimmingSupport-04618  
  [VkDisplayNativeHdrSurfaceCapabilitiesAMD](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayNativeHdrSurfaceCapabilitiesAMD.html)::`localDimmingSupport`
  **must** be supported

## <a href="#_see_also" class="anchor"></a>See Also

[VK_AMD_display_native_hdr](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_AMD_display_native_hdr.html),
[VkBool32](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkBool32.html), [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html),
[VkSwapchainKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSwapchainKHR.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkSetLocalDimmingAMD"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
