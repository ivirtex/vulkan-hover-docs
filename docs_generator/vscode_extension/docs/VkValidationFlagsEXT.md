# VkValidationFlagsEXT(3) Manual Page

## Name

VkValidationFlagsEXT - Specify validation checks to disable for a Vulkan instance



## [](#_c_specification)C Specification

When creating a Vulkan instance for which you wish to disable validation checks, add a [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationFlagsEXT.html) structure to the `pNext` chain of the [VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/latest/man/html/VkInstanceCreateInfo.html) structure, specifying the checks to be disabled.

```c++
// Provided by VK_EXT_validation_flags
typedef struct VkValidationFlagsEXT {
    VkStructureType                sType;
    const void*                    pNext;
    uint32_t                       disabledValidationCheckCount;
    const VkValidationCheckEXT*    pDisabledValidationChecks;
} VkValidationFlagsEXT;
```

## [](#_members)Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html) value identifying this structure.
- `pNext` is `NULL` or a pointer to a structure extending this structure.
- `disabledValidationCheckCount` is the number of checks to disable.
- `pDisabledValidationChecks` is a pointer to an array of [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCheckEXT.html) values specifying the validation checks to be disabled.

## [](#_description)Description

Valid Usage (Implicit)

- [](#VUID-VkValidationFlagsEXT-sType-sType)VUID-VkValidationFlagsEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VALIDATION_FLAGS_EXT`
- [](#VUID-VkValidationFlagsEXT-pDisabledValidationChecks-parameter)VUID-VkValidationFlagsEXT-pDisabledValidationChecks-parameter  
  `pDisabledValidationChecks` **must** be a valid pointer to an array of `disabledValidationCheckCount` valid [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCheckEXT.html) values
- [](#VUID-VkValidationFlagsEXT-disabledValidationCheckCount-arraylength)VUID-VkValidationFlagsEXT-disabledValidationCheckCount-arraylength  
  `disabledValidationCheckCount` **must** be greater than `0`

## [](#_see_also)See Also

[VK\_EXT\_validation\_flags](https://registry.khronos.org/vulkan/specs/latest/man/html/VK_EXT_validation_flags.html), [VkStructureType](https://registry.khronos.org/vulkan/specs/latest/man/html/VkStructureType.html), [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/latest/man/html/VkValidationCheckEXT.html)

## [](#_document_notes)Document Notes

For more information, see the [Vulkan Specification](https://registry.khronos.org/vulkan/specs/latest/html/vkspec.html#VkValidationFlagsEXT)

This page is extracted from the Vulkan Specification. Fixes and changes should be made to the Specification, not directly.

## [](#_copyright)Copyright

Copyright 2014-2025 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0