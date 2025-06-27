# VkCopyMicromapModeEXT(3) Manual Page

## Name

VkCopyMicromapModeEXT - Micromap copy mode



## <a href="#_c_specification" class="anchor"></a>C Specification

Possible values of `mode` specifying additional operations to perform
during the copy, are:

``` c
// Provided by VK_EXT_opacity_micromap
typedef enum VkCopyMicromapModeEXT {
    VK_COPY_MICROMAP_MODE_CLONE_EXT = 0,
    VK_COPY_MICROMAP_MODE_SERIALIZE_EXT = 1,
    VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT = 2,
    VK_COPY_MICROMAP_MODE_COMPACT_EXT = 3,
} VkCopyMicromapModeEXT;
```

## <a href="#_description" class="anchor"></a>Description

- `VK_COPY_MICROMAP_MODE_CLONE_EXT` creates a direct copy of the
  micromap specified in `src` into the one specified by `dst`. The `dst`
  micromap **must** have been created with the same parameters as `src`.

- `VK_COPY_MICROMAP_MODE_SERIALIZE_EXT` serializes the micromap to a
  semi-opaque format which can be reloaded on a compatible
  implementation.

- `VK_COPY_MICROMAP_MODE_DESERIALIZE_EXT` deserializes the semi-opaque
  serialization format in the buffer to the micromap.

- `VK_COPY_MICROMAP_MODE_COMPACT_EXT` creates a more compact version of
  a micromap `src` into `dst`. The micromap `dst` **must** have been
  created with a size at least as large as that returned by
  [vkCmdWriteMicromapsPropertiesEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdWriteMicromapsPropertiesEXT.html)
  after the build of the micromap specified by `src`.

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_opacity_micromap](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_opacity_micromap.html),
[VkCopyMemoryToMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMemoryToMicromapInfoEXT.html),
[VkCopyMicromapInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapInfoEXT.html),
[VkCopyMicromapToMemoryInfoEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCopyMicromapToMemoryInfoEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkCopyMicromapModeEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
