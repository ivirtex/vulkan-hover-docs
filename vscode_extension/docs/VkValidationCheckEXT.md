# VkValidationCheckEXT(3) Manual Page

## Name

VkValidationCheckEXT - Specify validation checks to disable



## [](#_c_specification)C Specification

Possible values of elements of the [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFlagsEXT.html)::`pDisabledValidationChecks` array, specifying validation checks to be disabled, are:

```c++
// Provided by VK_EXT_validation_flags
typedef enum VkValidationCheckEXT {
    VK_VALIDATION_CHECK_ALL_EXT = 0,
    VK_VALIDATION_CHECK_SHADERS_EXT = 1,
} VkValidationCheckEXT;
```

## [](#_description)Description

- `VK_VALIDATION_CHECK_ALL_EXT` specifies that all validation checks are disabled.
- `VK_VALIDATION_CHECK_SHADERS_EXT` specifies that shader validation is disabled.

## [](#_see_also)See Also

[VK\_EXT\_validation\_flags](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_flags.html), [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFlagsEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationCheckEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0