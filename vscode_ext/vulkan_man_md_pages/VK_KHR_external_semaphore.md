# VK_KHR_external_semaphore(3) Manual Page

## Name

VK_KHR_external_semaphore - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

78

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

[VK_KHR_external_semaphore_capabilities](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_KHR_external_semaphore_capabilities.html)  

## <a href="#_deprecation_state" class="anchor"></a>Deprecation State

- *Promoted* to <a
  href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#versions-1.1-promotions"
  target="_blank" rel="noopener">Vulkan 1.1</a>

## <a href="#_contact" class="anchor"></a>Contact

- James Jones <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_KHR_external_semaphore%5D%20@cubanismo%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_KHR_external_semaphore%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>cubanismo</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2016-10-21

**IP Status**  
No known IP claims.

**Contributors**  
- Faith Ekstrand, Intel

- Jesse Hall, Google

- Tobias Hector, Imagination Technologies

- James Jones, NVIDIA

- Jeff Juliano, NVIDIA

- Matthew Netsch, Qualcomm Technologies, Inc.

- Ray Smith, ARM

- Lina Versace, Google

## <a href="#_description" class="anchor"></a>Description

An application using external memory may wish to synchronize access to
that memory using semaphores. This extension enables an application to
create semaphores from which non-Vulkan handles that reference the
underlying synchronization primitive can be exported.

## <a href="#_promotion_to_vulkan_1_1" class="anchor"></a>Promotion to Vulkan 1.1

All functionality in this extension is included in core Vulkan 1.1, with
the KHR suffix omitted. The original type, enum and command names are
still available as aliases of the core functionality.

## <a href="#_new_structures" class="anchor"></a>New Structures

- Extending [VkSemaphoreCreateInfo](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreCreateInfo.html):

  - [VkExportSemaphoreCreateInfoKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkExportSemaphoreCreateInfoKHR.html)

## <a href="#_new_enums" class="anchor"></a>New Enums

- [VkSemaphoreImportFlagBitsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBitsKHR.html)

## <a href="#_new_bitmasks" class="anchor"></a>New Bitmasks

- [VkSemaphoreImportFlagsKHR](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagsKHR.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_KHR_EXTERNAL_SEMAPHORE_EXTENSION_NAME`

- `VK_KHR_EXTERNAL_SEMAPHORE_SPEC_VERSION`

- Extending [VkSemaphoreImportFlagBits](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkSemaphoreImportFlagBits.html):

  - `VK_SEMAPHORE_IMPORT_TEMPORARY_BIT_KHR`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_EXPORT_SEMAPHORE_CREATE_INFO_KHR`

## <a href="#_issues" class="anchor"></a>Issues

1\) Should there be restrictions on what side effects can occur when
waiting on imported semaphores that are in an invalid state?

**RESOLVED**: Yes. Normally, validating such state would be the
responsibility of the application, and the implementation would be free
to enter an undefined state if valid usage rules were violated. However,
this could cause security concerns when using imported semaphores, as it
would require the importing application to trust the exporting
application to ensure the state is valid. Requiring this level of trust
is undesirable for many potential use cases.

2\) Must implementations validate external handles the application
provides as input to semaphore state import operations?

**RESOLVED**: Implementations must return an error to the application if
the provided semaphore state handle cannot be used to complete the
requested import operation. However, implementations need not validate
handles are of the exact type specified by the application.

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2016-10-21 (James Jones)

  - Initial revision

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_KHR_external_semaphore"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
