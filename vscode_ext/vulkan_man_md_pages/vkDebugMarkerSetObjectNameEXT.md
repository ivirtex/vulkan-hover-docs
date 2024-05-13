# vkDebugMarkerSetObjectNameEXT(3) Manual Page

## Name

vkDebugMarkerSetObjectNameEXT - Give a user-friendly name to an object



## <a href="#_c_specification" class="anchor"></a>C Specification

An object can be given a user-friendly name by calling:

``` c
// Provided by VK_EXT_debug_marker
VkResult vkDebugMarkerSetObjectNameEXT(
    VkDevice                                    device,
    const VkDebugMarkerObjectNameInfoEXT*       pNameInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the object.

- `pNameInfo` is a pointer to a
  [VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectNameInfoEXT.html)
  structure specifying the parameters of the name to set on the object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDebugMarkerSetObjectNameEXT-device-parameter"
  id="VUID-vkDebugMarkerSetObjectNameEXT-device-parameter"></a>
  VUID-vkDebugMarkerSetObjectNameEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDebugMarkerSetObjectNameEXT-pNameInfo-parameter"
  id="VUID-vkDebugMarkerSetObjectNameEXT-pNameInfo-parameter"></a>
  VUID-vkDebugMarkerSetObjectNameEXT-pNameInfo-parameter  
  `pNameInfo` **must** be a valid pointer to a valid
  [VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectNameInfoEXT.html)
  structure

Host Synchronization

- Host access to `pNameInfo->object` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VkDebugMarkerObjectNameInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectNameInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDebugMarkerSetObjectNameEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
