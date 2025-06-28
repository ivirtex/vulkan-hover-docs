# VkExportMetalObjectTypeFlagBitsEXT(3) Manual Page

## Name

VkExportMetalObjectTypeFlagBitsEXT - Bitmask specifying Metal object types that can be exported from a Vulkan object



## [](#_c_specification)C Specification

Bits which indicate the types of Metal objects that may be exported from a corresponding Vulkan object are:

```c++
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

## [](#_description)Description

- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_DEVICE_BIT_EXT` specifies that a Metal `MTLDevice` may be exported.
- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_COMMAND_QUEUE_BIT_EXT` specifies that a Metal `MTLCommandQueue` may be exported.
- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_BUFFER_BIT_EXT` specifies that a Metal `MTLBuffer` may be exported.
- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_TEXTURE_BIT_EXT` specifies that a Metal `MTLTexture` may be exported.
- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_IOSURFACE_BIT_EXT` specifies that a Metal `IOSurface` may be exported.
- `VK_EXPORT_METAL_OBJECT_TYPE_METAL_SHARED_EVENT_BIT_EXT` specifies that a Metal `MTLSharedEvent` may be exported.

## [](#_see_also)See Also

[VK\_EXT\_metal\_objects](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_metal_objects.html), [VkExportMetalObjectCreateInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectCreateInfoEXT.html), [VkExportMetalObjectTypeFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkExportMetalObjectTypeFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkExportMetalObjectTypeFlagBitsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0