# vkDisplayPowerControlEXT(3) Manual Page

## Name

vkDisplayPowerControlEXT - Set the power state of a display



## <a href="#_c_specification" class="anchor"></a>C Specification

To set the power state of a display, call:

``` c
// Provided by VK_EXT_display_control
VkResult vkDisplayPowerControlEXT(
    VkDevice                                    device,
    VkDisplayKHR                                display,
    const VkDisplayPowerInfoEXT*                pDisplayPowerInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is a logical device associated with `display`.

- `display` is the display whose power state is modified.

- `pDisplayPowerInfo` is a pointer to a
  [VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerInfoEXT.html) structure
  specifying the new power state of `display`.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDisplayPowerControlEXT-device-parameter"
  id="VUID-vkDisplayPowerControlEXT-device-parameter"></a>
  VUID-vkDisplayPowerControlEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDisplayPowerControlEXT-display-parameter"
  id="VUID-vkDisplayPowerControlEXT-display-parameter"></a>
  VUID-vkDisplayPowerControlEXT-display-parameter  
  `display` **must** be a valid [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html) handle

- <a href="#VUID-vkDisplayPowerControlEXT-pDisplayPowerInfo-parameter"
  id="VUID-vkDisplayPowerControlEXT-pDisplayPowerInfo-parameter"></a>
  VUID-vkDisplayPowerControlEXT-pDisplayPowerInfo-parameter  
  `pDisplayPowerInfo` **must** be a valid pointer to a valid
  [VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerInfoEXT.html) structure

- <a href="#VUID-vkDisplayPowerControlEXT-commonparent"
  id="VUID-vkDisplayPowerControlEXT-commonparent"></a>
  VUID-vkDisplayPowerControlEXT-commonparent  
  Both of `device`, and `display` **must** have been created, allocated,
  or retrieved from the same [VkPhysicalDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkPhysicalDevice.html)

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_display_control](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_display_control.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html), [VkDisplayKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayKHR.html),
[VkDisplayPowerInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDisplayPowerInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDisplayPowerControlEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
