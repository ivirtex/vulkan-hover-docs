# VkValidationFlagsEXT(3) Manual Page

## Name

VkValidationFlagsEXT - Specify validation checks to disable for a Vulkan
instance



## <a href="#_c_specification" class="anchor"></a>C Specification

When creating a Vulkan instance for which you wish to disable validation
checks, add a [VkValidationFlagsEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationFlagsEXT.html)
structure to the `pNext` chain of the
[VkInstanceCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkInstanceCreateInfo.html) structure, specifying
the checks to be disabled.

``` c
// Provided by VK_EXT_validation_flags
typedef struct VkValidationFlagsEXT {
    VkStructureType                sType;
    const void*                    pNext;
    uint32_t                       disabledValidationCheckCount;
    const VkValidationCheckEXT*    pDisabledValidationChecks;
} VkValidationFlagsEXT;
```

## <a href="#_members" class="anchor"></a>Members

- `sType` is a [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html) value identifying
  this structure.

- `pNext` is `NULL` or a pointer to a structure extending this
  structure.

- `disabledValidationCheckCount` is the number of checks to disable.

- `pDisabledValidationChecks` is a pointer to an array of
  [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCheckEXT.html) values specifying
  the validation checks to be disabled.

## <a href="#_description" class="anchor"></a>Description

Valid Usage (Implicit)

- <a href="#VUID-VkValidationFlagsEXT-sType-sType"
  id="VUID-VkValidationFlagsEXT-sType-sType"></a>
  VUID-VkValidationFlagsEXT-sType-sType  
  `sType` **must** be `VK_STRUCTURE_TYPE_VALIDATION_FLAGS_EXT`

- <a href="#VUID-VkValidationFlagsEXT-pDisabledValidationChecks-parameter"
  id="VUID-VkValidationFlagsEXT-pDisabledValidationChecks-parameter"></a>
  VUID-VkValidationFlagsEXT-pDisabledValidationChecks-parameter  
  `pDisabledValidationChecks` **must** be a valid pointer to an array of
  `disabledValidationCheckCount` valid
  [VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCheckEXT.html) values

- <a
  href="#VUID-VkValidationFlagsEXT-disabledValidationCheckCount-arraylength"
  id="VUID-VkValidationFlagsEXT-disabledValidationCheckCount-arraylength"></a>
  VUID-VkValidationFlagsEXT-disabledValidationCheckCount-arraylength  
  `disabledValidationCheckCount` **must** be greater than `0`

## <a href="#_see_also" class="anchor"></a>See Also

[VK_EXT_validation_flags](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_validation_flags.html),
[VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html),
[VkValidationCheckEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkValidationCheckEXT.html)

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VkValidationFlagsEXT"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is extracted from the Vulkan Specification. Fixes and changes
should be made to the Specification, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.290  
Last updated 2024-07-11 23:39:16 -0700
