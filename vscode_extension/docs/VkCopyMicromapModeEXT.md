# VkCopyMicromapModeEXT(3) Manual Page

## Name

VkCopyMicromapModeEXT - Micromap copy mode



## [](#_c_specification)C Specification

Possible values of `mode` specifying additional operations to perform during the copy, are:

```c++
// Provided by VK_EXT_opacity_micromap
typedef enum VkCopyMicromapModeEXT {
    VK_COPY_MICROMAP_MODE_CLONE_EXT = 0,
    VK_COPY_MICROMAP_MODE_SERIALIZE_EXT = 1,
    VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT = 2,
    VK_COPY_MICROMAP_MODE_COMPACT_EXT = 3,
} VkCopyMicromapModeEXT;
```

## [](#_description)Description

- `VK_COPY_MICROMAP_MODE_CLONE_EXT` creates a direct copy of the micromap specified in `src` into the one specified by `dst`. The `dst` micromap **must** have been created with the same parameters as `src`.
- `VK_COPY_MICROMAP_MODE_SERIALIZE_EXT` serializes the micromap to a semi-opaque format which can be reloaded on a compatible implementation.
- `VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT` deserializes the semi-opaque serialization format in the buffer to the micromap.
- `VK_COPY_MICROMAP_MODE_COMPACT_EXT` creates a more compact version of a micromap `src` into `dst`. The micromap `dst` **must** have been created with a size at least as large as that returned by [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/vkCmdWriteMicromapsPropertiesEXT.html) after the build of the micromap specified by `src`.

## [](#_see_also)See Also

[VK\_EXT\_opacity\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_opacity_micromap.html), [VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMemoryToMicromapInfoEXT.html), [VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapInfoEXT.html), [VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkCopyMicromapToMemoryInfoEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkCopyMicromapModeEXT).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0