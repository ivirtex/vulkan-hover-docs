# vkDebugMarkerSetObjectTagEXT(3) Manual Page

## Name

vkDebugMarkerSetObjectTagEXT - Attach arbitrary data to an object



## <a href="#_c_specification" class="anchor"></a>C Specification

In addition to setting a name for an object, debugging and validation
layers may have uses for additional binary data on a per-object basis
that has no other place in the Vulkan API. For example, a
`VkShaderModule` could have additional debugging data attached to it to
aid in offline shader tracing. To attach data to an object, call:

``` c
// Provided by VK_EXT_debug_marker
VkResult vkDebugMarkerSetObjectTagEXT(
    VkDevice                                    device,
    const VkDebugMarkerObjectTagInfoEXT*        pTagInfo);
```

## <a href="#_parameters" class="anchor"></a>Parameters

- `device` is the device that created the object.

- `pTagInfo` is a pointer to a
  [VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectTagInfoEXT.html)
  structure specifying the parameters of the tag to attach to the
  object.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-vkDebugMarkerSetObjectTagEXT-device-parameter"
  id="VUID-vkDebugMarkerSetObjectTagEXT-device-parameter"></a>
  VUID-vkDebugMarkerSetObjectTagEXT-device-parameter  
  `device` **must** be a valid [VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html) handle

- <a href="#VUID-vkDebugMarkerSetObjectTagEXT-pTagInfo-parameter"
  id="VUID-vkDebugMarkerSetObjectTagEXT-pTagInfo-parameter"></a>
  VUID-vkDebugMarkerSetObjectTagEXT-pTagInfo-parameter  
  `pTagInfo` **must** be a valid pointer to a valid
  [VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectTagInfoEXT.html)
  structure

Host Synchronization

- Host access to `pTagInfo->object` **must** be externally synchronized

Return Codes

On success, this command returns  
- `VK_SUCCESS`

On failure, this command returns  
- `VK_ERROR_OUT_OF_HOST_MEMORY`

- `VK_ERROR_OUT_OF_DEVICE_MEMORY`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_debug_marker](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_marker.html),
[VkDebugMarkerObjectTagInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugMarkerObjectTagInfoEXT.html),
[VkDevice](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDevice.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#vkDebugMarkerSetObjectTagEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
