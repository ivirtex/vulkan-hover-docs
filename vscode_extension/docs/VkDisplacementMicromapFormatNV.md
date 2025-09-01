# VkDisplacementMicromapFormatNV(3) Manual Page

## Name

VkDisplacementMicromapFormatNV - Format enum for displacement micromaps



## [](#_c_specification)C Specification

Formats which **can** be set in [VkMicromapUsageEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapUsageEXT.html)::`format` and [VkMicromapTriangleEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkMicromapTriangleEXT.html)::`format` for micromap builds, are:

```c++
// Provided by VK_NV_displacement_micromap
typedef enum VkDisplacementMicromapFormatNV {
    VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV = 1,
    VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV = 2,
    VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV = 3,
} VkDisplacementMicromapFormatNV;
```

## [](#_description)Description

- `VK_DISPLACEMENT_MICROMAP_FORMAT_64_TRIANGLES_64_BYTES_NV` specifies that the given micromap format encodes 64 micro-triangles worth of displacements in 64 bytes as described in [Displacement Micromap Encoding](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#displacement-micromap-encoding).
- `VK_DISPLACEMENT_MICROMAP_FORMAT_256_TRIANGLES_128_BYTES_NV` specifies that the given micromap format encodes 256 micro-triangles worth of displacements in 128 bytes as described in [Displacement Micromap Encoding](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#displacement-micromap-encoding).
- `VK_DISPLACEMENT_MICROMAP_FORMAT_1024_TRIANGLES_128_BYTES_NV` specifies that the given micromap format encodes 1024 micro-triangles worth of displacements in 128 bytes as described in [Displacement Micromap Encoding](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#displacement-micromap-encoding).

Note

For compactness, these values are stored as 16-bit in some structures.

## [](#_see_also)See Also

[VK\_NV\_displacement\_micromap](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_NV_displacement_micromap.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkDisplacementMicromapFormatNV).

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0