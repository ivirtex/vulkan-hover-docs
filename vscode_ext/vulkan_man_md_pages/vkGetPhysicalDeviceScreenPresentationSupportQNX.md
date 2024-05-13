# vkGetPhysicalDeviceScreenPresentationSupportQNX(3) Manual Page

## Name

vkGetPhysicalDeviceScreenPresentationSupportQNX - Query physical device
for presentation to QNX Screen



## <a href="#_c_specification" class="anchor"></a>C Specification

To determine whether a queue family of a physical device supports
presentation to a QNX Screen compositor, call:

``` c
// Provided by VK_QNX_screen_surface
VkBool32 vkGetPhysicalDeviceScreenPresentationSupportQNX(
    VkPhysicalDevice                            physicalDevice,
    uint32_t                                    queueFamilyIndex,
    struct _screen_window*                      window);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `physicalDevice` is the physical device.

- `queueFamilyIndex` is the queue family index.

- `window` is the QNX Screen `window` object.

## <a href="#_description" class="anchor"></a>Description

This platform-specific function **can** be called prior to creating a
surface.

Valid Usage

- <a
  href="#VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-queueFamilyIndex-04743"
  id="VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-queueFamilyIndex-04743"></a>
  VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-queueFamilyIndex-04743  
  `queueFamilyIndex` **must** be less than `pQueueFamilyPropertyCount`
  returned by `vkGetPhysicalDeviceQueueFamilyProperties` for the given
  `physicalDevice`

Valid Usage (Implicit)

- <a
  href="#VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-physicalDevice-parameter"
  id="VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-physicalDevice-parameter"></a>
  VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-physicalDevice-parameter  
  `physicalDevice` **must** be a valid
  [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html) handle

- <a
  href="#VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-window-parameter"
  id="VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-window-parameter"></a>
  VUID-vkGetPhysicalDeviceScreenPresentationSupportQNX-window-parameter  
  `window` **must** be a valid pointer to a `_screen_window` value

## <a href="#_see_also" class="anchor"></a>See Also

[VK_QNX_screen_surface](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_QNX_screen_surface.html),
[VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkGetPhysicalDeviceScreenPresentationSupportQNX"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
