# VkExportMetalObjectTypeFlagBitsEXT(3) Manual Page

## Name

VkExportMetalObjectTypeFlagBitsEXT - Bitmask specifying Metal object
types that can be exported from a Vulkan object



## <a href="#_c_specification" class="anchor"></a>C Specification

Bits which indicate the types of Metal objects that may be exported from
a corresponding Vulkan object are:

``` c
// Provided by VK_EXT_metal_objects
typedef enum VkExportMetalObjectTypeFlagBitsEXT {
    VK_EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT = 0x00000001,
    VK_EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT = 0x00000002,
    VK_EXPORT_METAL_OBJECT_TYPE_METAL_BUFFER_BIT_EXT = 0x00000004,
    VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT = 0x00000008,
    VK_EXPORT_METAL_OBJECT_TYPE_METAL_IOSURFACE_BIT_EXT = 0x00000010,
    VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT = 0x00000020,
} VkExportMetalObjectTypeFlagBitsEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT` indicates a Metal
  `MTLDevice` may be exported.

- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT` indicates a
  Metal `MTLCommandQueue` may be exported.

- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_BUFFER_BIT_EXT` indicates a Metal
  `MTLBuffer` may be exported.

- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` indicates a Metal
  `MTLTexture` may be exported.

- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_IOSURFACE_BIT_EXT` indicates a
  Metal `IOSurface` may be exported.

- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT` indicates a
  Metal `MTLSharedEvent` may be exported.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_metal_objects](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_metal_objects.html),
[VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectCreateInfoEXT.html),
[VkExportMetalObjectTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportMetalObjectTypeFlagsEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkExportMetalObjectTypeFlagBitsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
