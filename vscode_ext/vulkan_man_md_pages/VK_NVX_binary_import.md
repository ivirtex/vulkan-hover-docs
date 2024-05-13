# VK_NVX_binary_import(3) Manual Page

## Name

VK_NVX_binary_import - device extension



## <a href="#_registered_extension_number" class="anchor"></a>Registered Extension Number

30

## <a href="#_revision" class="anchor"></a>Revision

1

## <a href="#_ratification_status" class="anchor"></a>Ratification Status

Not ratified

## <a href="#_extension_and_version_dependencies" class="anchor"></a>Extension and Version Dependencies

None

## <a href="#_api_interactions" class="anchor"></a>API Interactions

- Interacts with VK_EXT_debug_report

## <a href="#_contact" class="anchor"></a>Contact

- Eric Werness <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_binary_import%5D%20@ewerness-nv%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_binary_import%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>ewerness-nv</a>

- Liam Middlebrook <a
  href="https://github.com/KhronosGroup/Vulkan-Docs/issues/new?body=%5BVK_NVX_binary_import%5D%20@liam-middlebrook%0A*Here%20describe%20the%20issue%20or%20question%20you%20have%20about%20the%20VK_NVX_binary_import%20extension*"
  target="_blank" rel="nofollow noopener"><em></em>liam-middlebrook</a>

## <a href="#_other_extension_metadata" class="anchor"></a>Other Extension Metadata

**Last Modified Date**  
2021-04-09

**Contributors**  
- Eric Werness, NVIDIA

- Liam Middlebrook, NVIDIA

## <a href="#_description" class="anchor"></a>Description

This extension allows applications to import CuBIN binaries and execute
them.

<table>
<colgroup>
<col style="width: 50%" />
<col style="width: 50%" />
</colgroup>
<tbody>
<tr class="odd">
<td class="icon"><em></em></td>
<td class="content">Note
<p>There is currently no specification language written for this
extension. The links to APIs defined by the extension are to stubs that
only include generated content such as API declarations and implicit
valid usage statements.</p></td>
</tr>
</tbody>
</table>

## <a href="#_new_object_types" class="anchor"></a>New Object Types

- [VkCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuFunctionNVX.html)

- [VkCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleNVX.html)

## <a href="#_new_commands" class="anchor"></a>New Commands

- [vkCmdCuLaunchKernelNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCmdCuLaunchKernelNVX.html)

- [vkCreateCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCuFunctionNVX.html)

- [vkCreateCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkCreateCuModuleNVX.html)

- [vkDestroyCuFunctionNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCuFunctionNVX.html)

- [vkDestroyCuModuleNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/vkDestroyCuModuleNVX.html)

## <a href="#_new_structures" class="anchor"></a>New Structures

- [VkCuFunctionCreateInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuFunctionCreateInfoNVX.html)

- [VkCuLaunchInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuLaunchInfoNVX.html)

- [VkCuModuleCreateInfoNVX](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkCuModuleCreateInfoNVX.html)

## <a href="#_new_enum_constants" class="anchor"></a>New Enum Constants

- `VK_NVX_BINARY_IMPORT_EXTENSION_NAME`

- `VK_NVX_BINARY_IMPORT_SPEC_VERSION`

- Extending [VkObjectType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkObjectType.html):

  - `VK_OBJECT_TYPE_CU_FUNCTION_NVX`

  - `VK_OBJECT_TYPE_CU_MODULE_NVX`

- Extending [VkStructureType](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkStructureType.html):

  - `VK_STRUCTURE_TYPE_CU_FUNCTION_CREATE_INFO_NVX`

  - `VK_STRUCTURE_TYPE_CU_LAUNCH_INFO_NVX`

  - `VK_STRUCTURE_TYPE_CU_MODULE_CREATE_INFO_NVX`

If [VK_EXT_debug_report](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VK_EXT_debug_report.html) is supported:

- Extending
  [VkDebugReportObjectTypeEXT](https://registry.khronos.org/vulkan/specs/1.3-extensions/man/html/VkDebugReportObjectTypeEXT.html):

  - `VK_DEBUG_REPORT_OBJECT_TYPE_CU_FUNCTION_NVX_EXT`

  - `VK_DEBUG_REPORT_OBJECT_TYPE_CU_MODULE_NVX_EXT`

## <a href="#_version_history" class="anchor"></a>Version History

- Revision 1, 2021-04-09 (Eric Werness)

  - Internal revisions

## <a href="#_see_also" class="anchor"></a>See Also

No cross-references are available

## <a href="#_document_notes" class="anchor"></a>Document Notes

For more information, see the <a
href="https://registry.khronos.org/vulkan/specs/1.3-extensions/html/vkspec.html#VK_NVX_binary_import"
target="_blank" rel="noopener">Vulkan Specification</a>

This page is a generated document. Fixes and changes should be made to
the generator scripts, not directly.

## <a href="#_copyright" class="anchor"></a>Copyright

Copyright 2014-2024 The Khronos Group Inc.

SPDX-License-Identifier: CC-BY-4.0

Version 1.3.285  
Last updated 2024-05-10 01:10:25 -0700
